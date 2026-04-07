package resp

import (
	"strconv"
)

func decodeInteger(input string) (Value, error) {
	value, err := strconv.Atoi(input[1 : len(input)-len(terminator)])
	if err != nil {
		return EmptyValue(), err
	}
	return NewInt(value), nil
}
