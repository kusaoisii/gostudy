package trace

import (
	"io"
	"fmt"
)

// tracer はコードないでの出来事を記録できるオブジェクトを表すインタフェースです。
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t	*tracer) Trace(a ...interface{}){
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type nilTracer struct{}
func (t *nilTracer) Trace(a ...interface{}){}

//OffはTraceメソッドの呼び出しを無視するTracerを開始します
func Off() Tracer {
	return &nilTracer{}
}