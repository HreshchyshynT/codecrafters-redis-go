package resp

import (
	"bufio"
	"strconv"
)

func decodeInteger(r *bufio.Reader) (Value, error) {
	bytes, err := r.ReadBytes('\n')
	if err != nil {
		return EmptyValue(), err
	}
	value, err := strconv.Atoi(string(bytes[:len(bytes)-len(terminator)]))
	if err != nil {
		return EmptyValue(), err
	}
	return NewInt(value), nil
}
