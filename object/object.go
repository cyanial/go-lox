package object

const (
	String Type = iota
)

type Type uint8

type Object struct {
	Type Type

	value interface{}
}

//struct ObjString {
//	Obj obj;
//	int length;
//	char* chars;
//};

func NewString(s string) *Object {
	// alloc heap chars
	// mem copy

	return newObject(String, s)
}

func newObject(typ Type, value interface{}) *Object {
	return &Object{
		Type:  typ,
		value: value,
	}
}

func (o *Object) String() string {
	switch o.Type {
	case String:
		return o.AsString()

	default:
		return "unknown object"
	}
}

func (o *Object) AsString() string {
	return o.value.(string)
}

//func ObjType(v *value.Value) Type {
//	return v.AsObject().Type
//}

//func IsObjType(v *value.Value, typ Type) bool {
//	return v.Type == value.Object && v.AsObject().Type == typ
//}

//func ISString(v *value.Value) bool { return IsObjType(v, String) }

//func AsString(v *value.Value) *ObjectString {
//	return v.value.(*ObjectString)
//}

//func AsCString(v *value.Value) *ObjectString {
//	return v.value.(*ObjectString)->chars
//}
