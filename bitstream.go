/* dbitstream is a simple wrapper around a io.Reader and io.Writer to provide bit-level access to the stream. */
package dbitstream

import (
	"io"
)

type Bit bool

const (
	Zero Bit = false
	One      = true
)

type BitReader struct {
	r     io.Reader
	b     [1]byte
	count uint8
}

type BitWriter struct {
	w     io.Writer
	b     [1]byte
	count uint8
}

// NewReader returns a BitReader that returns a single bit at a time from 'r'
func NewReader(r io.Reader) *BitReader {
	b := new(BitReader)
	b.r = r
	return b
}

// ReadBit returns the next bit from the stream, reading a new byte from the underlying reader if required.
func (b *BitReader) ReadBit() (Bit, error) {
	if b.count == 0 {
		if n, err := b.r.Read(b.b[:]); n != 1 || err != nil {
			return Zero, err
		}
		b.count = 8
	}
	b.count--
	d := (b.b[0] & 0x80)
	b.b[0] <<= 1
	return d != 0, nil
}

// NewWriter returns a BitWriter that buffers bits and write the resulting bytes to 'w'
func NewWriter(w io.Writer) *BitWriter {
	b := new(BitWriter)
	b.w = w
	b.count = 8
	return b
}

// WriteBit writes a single bit to the stream, writing a new byte to 'w' if required.
func (b *BitWriter) WriteBit(bit Bit) error {

	if bit {
		b.b[0] |= 1 << (b.count - 1)
	}

	b.count--

	if b.count == 0 {
		if n, err := b.w.Write(b.b[:]); n != 1 || err != nil {
			return err
		}
		b.b[0] = 0
		b.count = 8
	}

	return nil
}

// WriteByte writes a single byte to the stream, regardless of alignment
func (b *BitWriter) WriteByte(byt byte) error {

	// fill up b.b with b.count bits from byt
	b.b[0] |= byt >> (8 - b.count)

	if n, err := b.w.Write(b.b[:]); n != 1 || err != nil {
		return err
	}
	b.b[0] = byt << b.count

	return nil
}

// ReadByte writes a single byte to the stream, regardless of alignment
func (b *BitReader) ReadByte() (byte, error) {

	if b.count == 0 {
		if n, err := b.r.Read(b.b[:]); n != 1 || err != nil {
			return 0, err
		}
		return b.b[0], nil
	}

	byt := b.b[0]

	if n, err := b.r.Read(b.b[:]); n != 1 || err != nil {
		return 0, err
	}

	byt |= b.b[0] >> b.count

	b.b[0] <<= (8 - b.count)

	return byt, nil
}

// Flush empties the currently in-process byte by filling it with 'bit'.
func (b *BitWriter) Flush(bit Bit) {

	for b.count != 8 {
		b.WriteBit(bit)
	}

	return
}
