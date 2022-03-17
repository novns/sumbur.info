// Code generated by qtc from "static.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package static

import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

func StreamStamp(qw422016 *qt422016.Writer, path string) {
	qw422016.N().S(`/static-stamp/`)
	qw422016.N().DL(TimeStamp(path))
	qw422016.N().S(`/`)
	qw422016.E().S(path)
}

func WriteStamp(qq422016 qtio422016.Writer, path string) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamStamp(qw422016, path)
	qt422016.ReleaseWriter(qw422016)
}

func Stamp(path string) string {
	qb422016 := qt422016.AcquireByteBuffer()
	WriteStamp(qb422016, path)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
