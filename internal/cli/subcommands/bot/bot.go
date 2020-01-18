package bot

import (
	"context"
	"time"

	"github.com/hortbot/hortbot/internal/bnsq"
	"github.com/hortbot/hortbot/internal/cli"
	"github.com/hortbot/hortbot/internal/cli/flags/botflags"
	"github.com/hortbot/hortbot/internal/cli/flags/jaegerflags"
	"github.com/hortbot/hortbot/internal/cli/flags/nsqflags"
	"github.com/hortbot/hortbot/internal/cli/flags/promflags"
	"github.com/hortbot/hortbot/internal/cli/flags/redisflags"
	"github.com/hortbot/hortbot/internal/cli/flags/sqlflags"
	"github.com/hortbot/hortbot/internal/cli/flags/twitchflags"
	"github.com/hortbot/hortbot/internal/pkg/ctxlog"
	"github.com/hortbot/hortbot/internal/pkg/errgroupx"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"golang.org/x/sync/semaphore"
)

const Name = "bot"

type config struct {
	cli.Common
	SQL        sqlflags.SQL
	Twitch     twitchflags.Twitch
	Redis      redisflags.Redis
	Bot        botflags.Bot
	NSQ        nsqflags.NSQ
	Jaeger     jaegerflags.Jaeger
	Prometheus promflags.Prometheus
}

func Run(args []string) {
	cli.Run(Name, args, &config{
		Common:     cli.DefaultCommon,
		SQL:        sqlflags.DefaultSQL,
		Twitch:     twitchflags.DefaultTwitch,
		Redis:      redisflags.DefaultRedis,
		Bot:        botflags.DefaultBot,
		NSQ:        nsqflags.DefaultNSQ,
		Jaeger:     jaegerflags.DefaultJaeger,
		Prometheus: promflags.Default,
	})
}

func (config *config) Main(ctx context.Context, _ []string) {
	defer config.Jaeger.Init(ctx, Name, config.Debug)()
	config.Prometheus.Run(ctx)

	driverName := config.SQL.DriverName()
	driverName = config.Jaeger.DriverName(ctx, driverName, config.Debug)
	db := config.SQL.Open(ctx, driverName)
	rdb := config.Redis.Client()
	twitchAPI := config.Twitch.Client()
	sender := config.NSQ.NewSendMessagePublisher()
	notifier := config.NSQ.NewNotifyPublisher()

	b := config.Bot.New(ctx, db, rdb, sender, notifier, twitchAPI)
	defer b.Stop()

	sem := semaphore.NewWeighted(int64(config.Bot.Workers))

	g := errgroupx.FromContext(ctx)

	incomingSub := config.NSQ.NewIncomingSubscriber(15*time.Second, func(i *bnsq.Incoming, metadata *bnsq.Metadata) error {
		ctx := metadata.With(ctx)
		ctx, span := trace.StartSpanWithRemoteParent(ctx, "OnIncoming", metadata.ParentSpan())
		defer span.End()

		if err := sem.Acquire(ctx, 1); err != nil {
			return err
		}

		g.Go(func(ctx context.Context) error {
			ctx, span := trace.StartSpanWithRemoteParent(ctx, "Worker", span.SpanContext())
			defer span.End()

			defer sem.Release(1)
			b.Handle(ctx, i.Origin, i.Message)
			return ctx.Err()
		})

		return nil
	})

	g.Go(sender.Run)
	g.Go(notifier.Run)
	g.Go(incomingSub.Run)

	if err := g.WaitIgnoreStop(); err != nil {
		ctxlog.Info(ctx, "exiting", zap.Error(err))
	}
}
