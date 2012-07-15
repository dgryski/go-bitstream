package dbitstream

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestStream(t *testing.T) {

	buf := bytes.NewBuffer(nil)
	br := NewReader(strings.NewReader("hello"))
	bw := NewWriter(buf)

	for {
		bit, err := br.GetBit()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Error("GetBit returned error err=", err.Error())
			return
		}
		bw.PutBit(bit)
	}

	s := buf.String()

	if s != "hello" {
		t.Error("got s=%s expected 'hello'", s)
	}
}
