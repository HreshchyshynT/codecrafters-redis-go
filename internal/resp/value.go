package resp

import "slices"

type Type string

const (
	BulkString Type = "BulkString"
	String     Type = "string"
	Array      Type = "array"
	Integer    Type = "integer"
	nothing    Type = "nothing"
)

type Value struct {
	Typ     Type
	String  string
	Array   []Value
	Integer int
}

func EmptyValue() Value {
	return Value{
		Typ: nothing,
	}
}

func (v *Value) isEmpty() bool {
	return v.Typ == nothing
}

func (v Value) Equals(other Value) bool {
	if v.Typ != other.Typ {
		return false
	}

	switch v.Typ {
	case nothing:
		return true
	case BulkString:
		fallthrough
	case String:
		return v.String == other.String
	case Array:
		if len(v.Array) != len(other.Array) {
			return false
		}
		return slices.EqualFunc(v.Array, other.Array, func(l, r Value) bool {
			return l.Equals(r)
		})
	case Integer:
		return v.Integer == other.Integer
	}
	return false
}

func NewBulkString(value string) Value {
	return Value{
		Typ:    BulkString,
		String: value,
	}
}

func NewString(value string) Value {
	return Value{
		Typ:    String,
		String: value,
	}
}

func NewArray(value []Value) Value {
	return Value{
		Typ:   Array,
		Array: value,
	}
}

func NewInt(value int) Value {
	return Value{Typ: Integer, Integer: value}
}
