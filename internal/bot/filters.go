package bot

import (
	"context"
	"database/sql"
	"net/url"

	"github.com/hortbot/hortbot/internal/db/models"
	"github.com/hortbot/hortbot/internal/linkmatch"
)

func tryFilter(ctx context.Context, s *Session) (filtered bool, err error) {
	if !s.Channel.ShouldModerate || !s.Channel.EnableFilters {
		return false, nil
	}

	if s.Channel.FilterLinks {
		filtered, err = filterLinks(ctx, s)
		if filtered || err != nil {
			return filtered, err
		}
	}

	return false, nil
}

func filterLinks(ctx context.Context, s *Session) (filtered bool, err error) {
	links := s.Links()

	if len(links) == 0 {
		return false, nil
	}

	if s.UserLevel.CanAccess(LevelSubscriber) {
		return false, nil
	}

	if allLinksPermitted(s.Channel.PermittedLinks, links) {
		return false, nil
	}

	permit, err := models.LinkPermits(
		models.LinkPermitWhere.ChannelID.EQ(s.Channel.ID),
		models.LinkPermitWhere.Name.EQ(s.User),
	).One(ctx, s.Tx)

	switch err {
	case nil:
		if err := permit.Delete(ctx, s.Tx); err != nil {
			return false, err
		}

		if s.Clock.Now().Before(permit.ExpiresAt) {
			return false, s.Replyf("Link permitted. (%s)", s.UserDisplay)
		}
	case sql.ErrNoRows:
		// Fall through
	default:
		return false, err
	}

	if err := s.DeleteMessage(); err != nil {
		return true, err
	}

	return true, s.Replyf("%s, please ask a moderator before posting links.", s.UserDisplay)
}

func allLinksPermitted(permitted []string, urls []*url.URL) bool {
	// Fast path for single links.
	if len(urls) == 1 {
		u := urls[0]

		for _, pd := range permitted {
			if linkmatch.HostAndPath(pd, u) {
				return true
			}
		}

		return false
	}

	urls = append(urls[:0:0], urls...)

	for _, pd := range permitted {
		allNil := true

		for i, u := range urls {
			if u == nil {
				continue
			}

			allNil = false

			if linkmatch.HostAndPath(pd, u) {
				urls[i] = nil
			}
		}

		if allNil {
			return true
		}
	}

	for _, u := range urls {
		if u != nil {
			return false
		}
	}

	return true
}