package resp

import (
	"bufio"
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
		return EmptyValue(), fmt.Errorf("invalid data format: %q", first)
	}
}

func encode(value Value) (string, error) {
	var encoded string
	var err error
	switch value.Typ {
	case BulkString:
		encoded = encodeBulkString(value)
	case String:
		encoded = encodeString(value)
	case Array:
		encoded, err = encodeArray(value.Array)
	case Integer:
		encoded = encodeInteger(value)
	case nothing:
		encoded = ""
	case NullBulkString:
		encoded = "$-1\r\n"
	}
	return encoded, err
}
