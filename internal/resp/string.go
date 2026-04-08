package resp

import (
	"bufio"
	"errors"
	"fmt"
)

/*
Simple strings are encoded as a plus (+) character, followed by a string. The
string mustn't contain a CR (\r) or LF (\n) character and is terminated by CRLF
(i.e., \r\n).
*/

func decodeString(r *bufio.Reader) (Value, error) {
	bytes, err := r.ReadBytes('\n')
	if err != nil {
		return EmptyValue(), err
	}
	if terminator != string(bytes[len(bytes)-terminatorLen:]) {
		return EmptyValue(), errors.New("invalid string")
	}

	result := string(bytes[:len(bytes)-terminatorLen])

	return NewString(result), nil
}

func encodeString(value Value) string {
	return fmt.Sprint("+", value.String, terminator)
}
