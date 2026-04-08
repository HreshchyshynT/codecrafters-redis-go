package resp

import (
	"bufio"
	"errors"
)

var InvalidFormat = errors.New("invalid data format")

const (
	/*
		The \r\n (CRLF) is the protocol's terminator, which always separates its parts.
	*/
	terminator = "\r\n"
)

func Decode(r *bufio.Reader) (Value, error) {
	first, _, err := r.ReadRune()
	if err != nil {
		return EmptyValue(), err
	}

	switch first {
	case '$':
		return decodeBulkString(r)
	case '*':
		return decodeArray(r)
	case ':':
		return decodeInteger(r)
	default:
		return EmptyValue(), InvalidFormat
	}
}

func Encode(body string) ([]byte, error) {
	// TODO: implement
	return []byte{}, nil
}
