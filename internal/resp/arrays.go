package resp

import (
	"bufio"
	"fmt"
	"strconv"
)

/*
Arrays
Clients send commands to the Redis server as RESP arrays. Similarly, some
Redis commands that return collections of elements use arrays as their
replies. An example is the LRANGE command that returns elements of a list.

RESP Arrays' encoding uses the following format:

*<number-of-elements>\r\n<element-1>...<element-n>

An asterisk (*) as the first byte.
One or more decimal digits (0..9) as the number of elements in the array as an unsigned, base-10 value.
The CRLF terminator.
An additional RESP type for every element of the array.
*/

func decodeArray(r *bufio.Reader) (Value, error) {
	lenBytes, _ := r.ReadBytes('\n')

	length, err := strconv.Atoi(string(lenBytes[:len(lenBytes)-terminatorLen]))
	if err != nil {
		return EmptyValue(), err
	}

	result := make([]Value, length)

	for i := range length {
		value, err := decode(r)
		if err != nil {
			return EmptyValue(), fmt.Errorf("Error parsing array: %v", err.Error())
		}
		result[i] = value
	}

	return NewArray(result), nil
}

func encodeArray() string {
	return ""
}
