package common

const (
	OpConstant OpCode = iota
	OpConstantLong
	OpAdd
	OpSubtract
	OpMultiply
	OpDivide
	OpNegate
	OpReturn
)

type OpCode uint8
