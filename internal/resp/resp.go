package resp

import (
	"errors"
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

	switch input[0] {
	case '$':
		return decodeBulkString(input)
	// case '*':
	// 	return decodeArray(input)
	default:
		return EmptyValue(), InvalidFormat
	}
}

func Encode(body string) ([]byte, error) {
	return []byte{}, nil
}
