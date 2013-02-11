package dbitstream

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestBitStream(t *testing.T) {

	buf := bytes.NewBuffer(nil)
	br := NewReader(strings.NewReader("hello"))
	bw := NewWriter(buf)

	for {
		bit, err := br.ReadBit()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Error("GetBit returned error err=", err.Error())
			return
		}
		bw.WriteBit(bit)
	}

	s := buf.String()

	if s != "hello" {
		t.Error("expected 'hello', got=", []byte(s))
	}
}

func TestByteStream(t *testing.T) {

	buf := bytes.NewBuffer(nil)
	br := NewReader(strings.NewReader("hello"))
	bw := NewWriter(buf)

	for i := 0; i < 3; i++ {
		bit, err := br.ReadBit()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Error("GetBit returned error err=", err.Error())
			return
		}
		bw.WriteBit(bit)
	}

	for i := 0; i < 4; i++ {
		byt, err := br.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Error("GetByte returned error err=", err.Error())
			return
		}
		bw.WriteByte(byt)
	}

	for i := 0; i < 5; i++ {
		bit, err := br.ReadBit()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Error("GetBit returned error err=", err.Error())
			return
		}
		bw.WriteBit(bit)
	}

	s := buf.String()

	if s != "hello" {
		t.Error("expected 'hello', got=", []byte(s))
	}
}
