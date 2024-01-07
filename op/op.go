package op

const (
	Constant Code = iota
	ConstantLong
	Nil
	True
	False
	Equal
	Greater
	Less
	Add
	Subtract
	Multiply
	Divide
	Not
	Negate
	Return
)

type Code byte
