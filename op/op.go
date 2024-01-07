package op

const (
	Constant Code = iota
	ConstantLong
	Nil
	True
	False
	Add
	Subtract
	Multiply
	Divide
	Not
	Negate
	Return
)

type Code byte
