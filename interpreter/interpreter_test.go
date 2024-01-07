package interpreter

import (
	"github.com/cyanial/go-lox/chunk"
	"github.com/cyanial/go-lox/env"
	"github.com/cyanial/go-lox/op"
	"testing"
)

func TestInterpreter(t *testing.T) {
	ck := chunk.New()
	line := 1

	env.DebugTraceExecution = true
	env.DebugPrintCode = true

	ck.AddConstant(1.2, line)
	ck.AddConstant(3.3, line)
	ck.AddOp(op.Add, line)
	ck.AddConstant(1, line)
	ck.AddOp(op.Subtract, line)
	ck.AddConstant(5.1, line)
	ck.AddOp(op.Multiply, line)
	ck.AddConstant(2, line)
	ck.AddOp(op.Divide, line)
	ck.AddOp(op.Negate, line)

	ck.AddOp(op.Return, line)

	it := New(ck)
	it.Run()
}

func TestInterpreterConstantLong(t *testing.T) {
	ck := chunk.New()
	line := 1

	env.DebugTraceExecution = true
	env.DebugPrintCode = true

	ck.AddConstant(1, line)

	for i := 0; i < 11111; i++ {
		ck.AddConstant(1, line)
		ck.AddOp(op.Add, line)
	}

	ck.AddOp(op.Negate, line)
	ck.AddOp(op.Return, line)
	it := New(ck)
	it.Run()
}
