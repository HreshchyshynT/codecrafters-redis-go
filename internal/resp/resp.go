package resp

import (
	"errors"
)

// Arrays
// Clients send commands to the Redis server as RESP arrays. Similarly, some
// Redis commands that return collections of elements use arrays as their
// replies. An example is the LRANGE command that returns elements of a list.
//
// RESP Arrays' encoding uses the following format:
//
// *<number-of-elements>\r\n<element-1>...<element-n>
//
// An asterisk (*) as the first byte.
// One or more decimal digits (0..9) as the number of elements in the array as an unsigned, base-10 value.
// The CRLF terminator.
// An additional RESP type for every element of the array.

var InvalidFormat = errors.New("invalid data format")

const (
	/*
		The \r\n (CRLF) is the protocol's terminator, which always separates its parts.
	*/
	terminator = "\r\n"
)

func Decode(data []byte) (string, error) {
	input := string(data)

	switch input[0] {
	case '$':
		return decodeBulkString(input)
	default:
		return "", InvalidFormat
	}

	return "", nil
}

func Encode(body string) ([]byte, error) {
	return []byte{}, nil
}
