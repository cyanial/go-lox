package common

const (
	OpConstant OpCode = iota
	OpConstantLong
	OpReturn
)

type OpCode uint8
