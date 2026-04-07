package resp

import (
	"errors"
	"strings"
)

var InvalidFormat = errors.New("invalid data format")

const (
	/*
		The \r\n (CRLF) is the protocol's terminator, which always separates its parts.
	*/
	terminator = "\r\n"
)

func Decode(data []byte) (Value, error) {
	input := string(data)
	if !strings.HasSuffix(input, terminator) {
		return EmptyValue(), InvalidFormat
	}

	switch input[0] {
	case '$':
		return decodeBulkString(input)
	// case '*':
	// 	return decodeArray(input)
	case ':':
		return decodeInteger(input)
	default:
		return EmptyValue(), InvalidFormat
	}
}

func Encode(body string) ([]byte, error) {
	// TODO: implement
	return []byte{}, nil
}
