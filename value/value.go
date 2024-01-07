package value

import "fmt"

const (
	Bool Type = iota
	Nil
	Number
)

type Type uint8

type Value struct {
	Type Type

	value interface{}
}

func NewNumber(v float64) *Value {
	return newValue(Number, v)
}

func NewNil() *Value {
	return newValue(Nil, nil)

}

func NewBool(b bool) *Value {
	return newValue(Bool, b)
}

func newValue(typ Type, v interface{}) *Value {
	return &Value{
		Type:  typ,
		value: v,
	}
}

func (v *Value) String() string {
	switch v.Type {
	case Bool:
		return fmt.Sprint(v.AsBool())
	case Nil:
		return "nil"
	case Number:
		return fmt.Sprintf("%g", v.AsNumber())
	}

	return "unknown value"
}

func (v *Value) AsBool() bool {
	return v.value.(bool)
}

func (v *Value) AsNumber() float64 {
	return v.value.(float64)
}

func (v *Value) Equal(vv *Value) bool {
	if v.Type != vv.Type {
		return false
	}
	switch v.Type {
	case Bool:
		return v.AsBool() == vv.AsBool()
	case Nil:
		return true
	case Number:
		return v.AsNumber() == vv.AsNumber()
	default:
		return false
	}
}
