package resp

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
)

/*
Bulk strings
A bulk string represents a single binary string. The string can be of any
size, but by default, Redis limits it to 512 MB (see the proto-max-bulk-len
configuration directive).
RESP encodes bulk strings in the following way:

$<length>\r\n<data>\r\n

The dollar sign ($) as the first byte.
One or more decimal digits (0..9) as the string's length, in bytes, as an unsigned, base-10 value.
The CRLF terminator.
The data.
A final CRLF.
*/

func decodeBulkString(r *bufio.Reader) (Value, error) {
	lengthBytes, _ := r.ReadBytes('\n')
	if lengthBytes[len(lengthBytes)-2] != '\r' {
		return EmptyValue(), errors.New("Invalid bulk string: length must be separated with content by \\r\\n")
	}
	length, err := strconv.Atoi(string(lengthBytes[:len(lengthBytes)-2]))
	if err != nil {
		return EmptyValue(), err
	}
	if length == 0 || length == -1 {
		// TODO: consider proper handling of NULL strings (when length == -1)
		return EmptyValue(), nil
	}

	contentBytes := make([]byte, length)

	i, err := r.Read(contentBytes)
	if err != nil {
		return EmptyValue(), err
	}
	last, err := r.ReadString('\n')
	if err != nil {
		return EmptyValue(), err
	}

	if i != length || last != terminator {
		return EmptyValue(), errors.New("Invalid bulk string. $length != data length")
	}

	return NewString(string(contentBytes)), nil
}

func encodeBulkString(input Value) string {
	return fmt.Sprint("$", len(input.String), terminator, input.String, terminator)
}
