package vm

import (
	"github.com/cyanial/go-lox/compiler"
	"github.com/cyanial/go-lox/interpreter"
)

type VirtualMachine struct {
	It *interpreter.Interpreter
	Cp *compiler.Compiler
}

func New() *VirtualMachine {
	return &VirtualMachine{
		Cp: compiler.New(),
	}
}

func (v *VirtualMachine) Interpret(source string) interpreter.Result {
	v.It = interpreter.New(v.Cp.Compile(source))

	return v.It.Run()
}
