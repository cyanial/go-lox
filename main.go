package main

import (
	"github.com/cyanial/go-interpreters/vm"
)

func main() {

	source := `
// this is comment
var hello = "world";
var age = 12;

func echo(var msg) {
	return msg;
};

print 1.2;
`
	v := vm.New(nil, true, true)
	v.Interpret(source)

	//c := common.NewChunk()
	//
	////constant := c.AddConstant(3.3)
	////c.AddOp(common.OpConstantLong, 101)
	////c.AddOp(common.OpCode(constant&0x0f), 101)
	////c.AddOp(common.OpCode(constant>>8&0x0f), 101)
	////c.AddOp(common.OpCode(constant>>16&0x0f), 101)
	//
	//constant := c.AddConstant(1.2)
	//c.AddOp(common.OpConstant, 123)
	//c.AddOp(common.OpCode(constant), 123)
	//
	//constant = c.AddConstant(3.4)
	//c.AddOp(common.OpConstant, 123)
	//c.AddOp(common.OpCode(constant), 123)
	//
	//c.AddOp(common.OpAdd, 123)
	//
	//constant = c.AddConstant(5.6)
	//c.AddOp(common.OpConstant, 123)
	//c.AddOp(common.OpCode(constant), 123)
	//
	//c.AddOp(common.OpDivide, 123)
	//
	//c.AddOp(common.OpNegate, 123)
	//
	//c.AddOp(common.OpReturn, 123)
	////disassembly.DisAssembleChunk(c, "test chunk")
	//v := vm.New(c, false)
	////v.Interpret()
}
