package resp

import (
	"bufio"
	"encoding/json"
	"io"
)

type Decoder struct {
	r   *bufio.Reader
	buf []byte
}

func NewDecoder(r io.Reader) *Decoder {
	_ = json.Decoder{}
	return &Decoder{
		r: bufio.NewReader(r),
	}
}

func (dec *Decoder) Decode() (Value, error) {
	return decode(dec.r)
}
