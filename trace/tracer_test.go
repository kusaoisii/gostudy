package trace

import (
	"testing"
	"bytes"
)

func TestNew(t *testing.T){
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Newから戻り値がnilです")
	} else {
		tracer.Trace("こんにちは、traceパッケージ")
		if buf.String() != "こんにちは、traceパッケージ\n" {
			t.Errorf("'%s'という余った文字列が出力されました",buf.String())
		}
	}
}

func TestOff(t *testing.T) {
	var slientTracer Tracer = Off()
	silentTracer.Trace("データ")
}