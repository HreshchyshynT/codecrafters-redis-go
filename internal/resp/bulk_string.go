package resp

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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
TODO: consider proper handling of NULL strings
*/
func decodeBulkString(input string) (string, error) {
	splitted := strings.Split(input, terminator)
	lenghtString, _ := strings.CutPrefix(splitted[0], "$")
	length, err := strconv.Atoi(lenghtString)
	if err != nil {
		return "", err
	}
	if length == 0 || length == -1 {
		return "", nil
	}

	if length != len(splitted[1]) {
		// not sure we need this
		return "", errors.New("Invalid bulk string. $length != data length")
	}

	return splitted[1], nil
}

func encodeBulkString(input string) string {
	return fmt.Sprint("$", len(input), terminator, input, terminator)
}
