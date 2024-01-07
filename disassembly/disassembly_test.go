package disassembly

import (
	"github.com/cyanial/go-lox/chunk"
	"github.com/cyanial/go-lox/op"
	"github.com/cyanial/go-lox/value"
	"testing"
)

func TestDisassemble(t *testing.T) {
	ck := chunk.New()
	line := 1

	ck.AddConstant(value.NewNumber(1.2), line)
	ck.AddOp(op.Return, line)

	DisassembleChunk(ck, "TestDisassemble")
}

func TestDisassembleAddConstantLong(t *testing.T) {
	ck := chunk.New()
	line := 1

	for i := 0; i < 11111; i++ {
		ck.AddConstant(value.NewNumber(1.2), line)
	}
	ck.AddOp(op.Return, line)

	DisassembleChunk(ck, "TestDisassembleAddConstantLong")
}

func TestDisassembleBoolNilNumber(t *testing.T) {
	ck := chunk.New()
	line := 1

	ck.AddConstant(value.NewNumber(123), line)
	ck.AddConstant(value.NewNil(), line)
	ck.AddConstant(value.NewBool(true), line)
	ck.AddConstant(value.NewBool(false), line)

	ck.AddOp(op.Return, line)

	DisassembleChunk(ck, "TestDisassembleBoolNilNumber")
}

func TestDisassembleLogicalNotAndFalseValue(t *testing.T) {
	ck := chunk.New()
	line := 1

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

	DisassembleChunk(ck, "TestDisassembleLogicalNotAndFalseValue")
}
