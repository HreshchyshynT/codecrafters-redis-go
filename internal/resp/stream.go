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

type Encoder struct {
	w io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		w: w,
	}
}

func (enc *Encoder) Encode(value Value) error {
	encoded, err := encode(value)
	if err != nil {
		return err
	}
	enc.w.Write([]byte(encoded))
	return nil
}
