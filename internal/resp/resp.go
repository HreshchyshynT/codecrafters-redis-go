package resp

import (
	"bufio"
	"errors"
	"fmt"
)

const (
	/*
		The \r\n (CRLF) is the protocol's terminator, which always separates its parts.
	*/
	terminator    = "\r\n"
	terminatorLen = 2
)

func decode(r *bufio.Reader) (Value, error) {
	first, err := r.ReadByte()
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
	case '+':
		return decodeString(r)
	default:
		return EmptyValue(), errors.New(fmt.Sprintf("invalid data format: %q", first))
	}
}

func encode(body string) ([]byte, error) {
	// TODO: implement
	return []byte{}, nil
}
