// Code generated by qtc from "channels.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package templates

import (
	"github.com/hortbot/hortbot/internal/db/models"

	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

type ChannelsPage struct {
	BasePage
	Channels models.ChannelSlice
}

func (p *ChannelsPage) StreamPageTitle(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
HortBot - Channels
`)
}

func (p *ChannelsPage) WritePageTitle(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageTitle(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *ChannelsPage) PageTitle() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageTitle(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (p *ChannelsPage) StreamPageBody(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
<section class="section">
    <div class="container">
        <h1 class="title has-text-centered">
            Channels
        </h1>

        <div class="columns">
            <div class="column is-8 is-offset-2">
                <div class="list">
                    `)
	for _, channel := range p.Channels {
		qw422016.N().S(`
                        <a class="list-item" href="/c/`)
		qw422016.N().U(channel.Name)
		qw422016.N().S(`">`)
		qw422016.E().S(channel.Name)
		qw422016.N().S(`</a>
                    `)
	}
	qw422016.N().S(`
                </div>
            </div>
        </div>
    </div>
</section>
`)
}

func (p *ChannelsPage) WritePageBody(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageBody(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *ChannelsPage) PageBody() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageBody(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}