package main

import (
	"github.com/cyanial/go-interpreters/common"
	"github.com/cyanial/go-interpreters/disassembly"
)

func main() {

	c := common.NewChunk()

	constant := c.AddConstant(3.3)
	c.AddOp(common.OpConstantLong, 101)
	c.AddOp(common.OpCode(constant&0x0f), 101)
	c.AddOp(common.OpCode(constant>>8&0x0f), 101)
	c.AddOp(common.OpCode(constant>>16&0x0f), 101)

	constant = c.AddConstant(1.2)
	c.AddOp(common.OpConstant, 123)
	c.AddOp(common.OpCode(constant), 123)

	c.AddOp(common.OpReturn, 123)
	disassembly.DisAssembleChunk(c, "test chunk")
}
