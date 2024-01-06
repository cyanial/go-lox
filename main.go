package main

import (
	"fmt"
	"github.com/cyanial/go-interpreters/compiler"
)

func main() {

	source := `
// this is comment
var hello = "world";
var age = 12;

func echo(var msg) {
	return msg;
};
`

	tokens := make([]*compiler.Token, 0)

	sc := compiler.NewScanner(source)
	for t := sc.ScanToken(); t.Type != compiler.TokenEOF; t = sc.ScanToken() {
		tokens = append(tokens, t)
		fmt.Printf("%#v\n", *t)
	}

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
