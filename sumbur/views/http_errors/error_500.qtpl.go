// Code generated by qtc from "error_500.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package http_errors

import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

func (data *Panic) StreamTitle(qw422016 *qt422016.Writer) {
	qw422016.N().S(`Ошибка 500`)
}

func (data *Panic) WriteTitle(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	data.StreamTitle(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (data *Panic) Title() string {
	qb422016 := qt422016.AcquireByteBuffer()
	data.WriteTitle(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (data *Panic) StreamBody(qw422016 *qt422016.Writer) {
	qw422016.N().S(`<p>Что-то пошло не так!

<pre>
`)
	qw422016.E().V(data.err)
	qw422016.N().S(`

`)
	qw422016.E().Z(data.stack)
	qw422016.N().S(`
</pre>
`)
}

func (data *Panic) WriteBody(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	data.StreamBody(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (data *Panic) Body() string {
	qb422016 := qt422016.AcquireByteBuffer()
	data.WriteBody(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}