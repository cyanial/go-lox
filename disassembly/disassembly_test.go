package disassembly

import (
	"github.com/cyanial/go-lox/chunk"
	"github.com/cyanial/go-lox/op"
	"testing"
)

func TestDisassemble(t *testing.T) {
	ck := chunk.New()
	line := 1

	ck.AddConstant(1.2, line)
	ck.AddOp(op.Return, line)

	DisassembleChunk(ck, "TestDisassemble")
}

func TestDisassembleAddConstantLong(t *testing.T) {
	ck := chunk.New()
	line := 1

	for i := 0; i < 11111; i++ {
		ck.AddConstant(1.2, line)
	}
	ck.AddOp(op.Return, line)

	DisassembleChunk(ck, "TestDisassembleAddConstantLong")
}
