package common

type Chunk struct {
	OpCodes   []OpCode
	Lines     []int
	Constants []Value
}

func NewChunk() *Chunk {
	return &Chunk{
		OpCodes:   make([]OpCode, 0),
		Constants: make([]Value, 0),
		Lines:     make([]int, 0),
	}
}

func (c *Chunk) AddOp(op OpCode, line int) {
	c.OpCodes = append(c.OpCodes, op)
	c.Lines = append(c.Lines, line)
}

func (c *Chunk) AddConstant(value Value) int {
	c.Constants = append(c.Constants, value)
	return len(c.Constants) - 1
}
