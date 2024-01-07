package interpreter

import (
	"github.com/cyanial/go-lox/chunk"
	"github.com/cyanial/go-lox/env"
	"github.com/cyanial/go-lox/op"
	"github.com/cyanial/go-lox/value"
	"testing"
)

func TestInterpreter(t *testing.T) {
	ck := chunk.New()
	line := 1

	env.DebugTraceExecution = true
	env.DebugPrintCode = true

	ck.AddConstant(value.NewNumber(1.2), line)
	ck.AddConstant(value.NewNumber(3.3), line)
	ck.AddOp(op.Add, line)
	ck.AddConstant(value.NewNumber(1), line)
	ck.AddOp(op.Subtract, line)
	ck.AddConstant(value.NewNumber(5.1), line)
	ck.AddOp(op.Multiply, line)
	ck.AddConstant(value.NewNumber(2), line)
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

	ck.AddConstant(value.NewNumber(1), line)

	for i := 0; i < 11111; i++ {
		ck.AddConstant(value.NewNumber(1), line)
		ck.AddOp(op.Add, line)
	}

	ck.AddOp(op.Negate, line)
	ck.AddOp(op.Return, line)
	it := New(ck)
	it.Run()
}

func TestInterpreterBoolNil(t *testing.T) {
	ck := chunk.New()
	line := 1

	env.DebugTraceExecution = true
	env.DebugPrintCode = true

	ck.AddOp(op.Nil, line)
	ck.AddOp(op.True, line)
	ck.AddOp(op.False, line)
	ck.AddOp(op.Return, line)

	it := New(ck)
	it.Run()
}

func TestInterpreterBinaryOpError(t *testing.T) {
	ck := chunk.New()
	line := 1

	env.DebugTraceExecution = true
	env.DebugPrintCode = true

	ck.AddOp(op.Nil, line)
	ck.AddOp(op.True, line)
	ck.AddOp(op.Divide, line)
	ck.AddOp(op.Return, line)

	it := New(ck)
	it.Run()
}

func TestInterpreterNegateError(t *testing.T) {
	ck := chunk.New()
	line := 1

	env.DebugTraceExecution = true
	env.DebugPrintCode = true

	ck.AddOp(op.Nil, line)
	ck.AddOp(op.Negate, line)
	ck.AddOp(op.Return, line)

	it := New(ck)
	it.Run()
}

func TestInterpreterLogicalNotAndFalseValue(t *testing.T) {
	ck := chunk.New()
	line := 1

	env.DebugTraceExecution = true
	env.DebugPrintCode = true

	ck.AddOp(op.Nil, line)
	ck.AddOp(op.Not, line)
	ck.AddOp(op.True, line)
	ck.AddOp(op.Not, line)
	ck.AddOp(op.False, line)
	ck.AddOp(op.Not, line)
	ck.AddConstant(value.NewNumber(12.1), line)
	ck.AddOp(op.Not, line)
	ck.AddConstant(value.NewNumber(0), line)
	ck.AddOp(op.Not, line)

	ck.AddOp(op.Return, line)

	it := New(ck)
	it.Run()
}
