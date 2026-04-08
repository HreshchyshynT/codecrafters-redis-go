package resp

import (
	"bufio"
	"errors"
	"fmt"
	"log"
)

const (
	/*
		The \r\n (CRLF) is the protocol's terminator, which always separates its parts.
	*/
	terminator    = "\r\n"
	terminatorLen = 2
)

func decode(r *bufio.Reader) (Value, error) {
	all, err := r.Peek(r.Size())
	log.Printf("all bytes: %q\n", string(all))

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
	default:
		return EmptyValue(), errors.New(fmt.Sprintf("invalid data format: %q", first))
	}
}

func encode(body string) ([]byte, error) {
	// TODO: implement
	return []byte{}, nil
}
