package chunk

import (
	"github.com/cyanial/go-lox/op"
	"github.com/cyanial/go-lox/value"
	"math"
)

type Chunk struct {
	Codes     []op.Code
	Constants []value.Value

	lines []int
}

func New() *Chunk {
	return &Chunk{
		Codes:     make([]op.Code, 0),
		Constants: make([]value.Value, 0),
		lines:     make([]int, 0),
	}
}

func (c *Chunk) AddOp(op op.Code, line int) {
	c.Codes = append(c.Codes, op)

	c.addLine(line)
}

func (c *Chunk) AddConstant(value value.Value, line int) {
	c.Constants = append(c.Constants, value)
	idx := len(c.Constants) - 1
	if idx <= math.MaxUint8 {
		c.AddOp(op.Constant, line)
		c.AddOp(op.Code(idx), line)
	} else {
		c.AddOp(op.ConstantLong, line)
		c.AddOp(op.Code(idx&0xff), line)
		c.AddOp(op.Code((idx>>8)&0xff), line)
		c.AddOp(op.Code((idx>>16)&0xff), line)
	}

	c.addLine(line)
}

func (c *Chunk) addLine(line int) {
	// Todo: refactor lines with run-length code
	c.lines = append(c.lines, line)
}

func (c *Chunk) GetLine(i int) int {
	// Todo: refactor lines with run-length code
	return c.lines[i]
}
