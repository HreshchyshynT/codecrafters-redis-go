package resp

import (
	"bufio"
	"fmt"
	"strconv"
)

/*
This type is a CRLF-terminated string that represents a signed, base-10, 64-bit integer.

RESP encodes integers in the following way:

:[<+|->]<value>\r\n

    The colon (:) as the first byte.
    An optional plus (+) or minus (-) as the sign.
    One or more decimal digits (0..9) as the integer's unsigned, base-10 value.
    The CRLF terminator.

For example, :0\r\n and :1000\r\n are integer replies (of zero and one thousand, respectively).
*/

func decodeInteger(r *bufio.Reader) (Value, error) {
	bytes, err := r.ReadBytes('\n')
	if err != nil {
		return EmptyValue(), err
	}
	value, err := strconv.Atoi(string(bytes[:len(bytes)-terminatorLen]))
	if err != nil {
		return EmptyValue(), err
	}
	return NewInt(value), nil
}

func encodeInteger(value Value) string {
	return fmt.Sprint(":", value.Integer, terminator)
}
