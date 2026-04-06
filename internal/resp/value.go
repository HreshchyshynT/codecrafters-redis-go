package resp

import "slices"

type Type string

const (
	String  Type = "string"
	Array   Type = "array"
	nothing Type = "nothing"
)

type Value struct {
	Typ    Type
	String string
	Array  []Value
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
	case String:
		return v.String == other.String
	case Array:
		if len(v.Array) != len(other.Array) {
			return false
		}
		return slices.EqualFunc(v.Array, other.Array, func(l, r Value) bool {
			return l.Equals(r)
		})
	}
	return false
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
