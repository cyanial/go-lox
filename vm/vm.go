package vm

import (
	"github.com/cyanial/go-lox/chunk"
	"github.com/cyanial/go-lox/interpreter"
)

type VirtualMachine struct {
	It *interpreter.Interpreter
}

func New(ck *chunk.Chunk) *VirtualMachine {
	return &VirtualMachine{
		It: interpreter.New(ck),
	}
}

func (v *VirtualMachine) Interpret(source string) interpreter.Result {
	// compile

	return v.It.Run()
}
