package op

const (
	Constant Code = iota
	ConstantLong
	Add
	Subtract
	Multiply
	Divide
	Negate
	Return
)

type Code byte
