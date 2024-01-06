package vm

import (
	"fmt"
	"github.com/cyanial/go-interpreters/common"
	"github.com/cyanial/go-interpreters/disassembly"
)

const (
	InterpretOK InterpretResult = iota
	InterpretCompileError
	InterpretRuntimeError
)

type InterpretResult int

type VM struct {
	Chunk *common.Chunk
	IP    int
	Stack []common.Value

	DebugTraceExecution bool
}

func New(c *common.Chunk, debugTraceExecution bool) *VM {
	return &VM{
		Chunk:               c,
		IP:                  0,
		Stack:               make([]common.Value, 0, 256),
		DebugTraceExecution: debugTraceExecution,
	}
}

func (vm *VM) Interpret() InterpretResult {
	return vm.run()
}

func (vm *VM) run() InterpretResult {

	for {
		if vm.DebugTraceExecution {
			fmt.Printf("      STACK: ")
			for _, slot := range vm.Stack {
				fmt.Printf("[%v]", slot)
			}
			fmt.Println()
			disassembly.DisassembleInstruction(vm.Chunk, vm.IP)
		}

		instruction := vm.readByte()
		switch instruction {
		case common.OpConstant:
			constant := vm.Chunk.Constants[vm.readByte()]
			vm.push(constant)
		case common.OpConstantLong:

		case common.OpAdd:
			fallthrough
		case common.OpSubtract:
			fallthrough
		case common.OpMultiply:
			fallthrough
		case common.OpDivide:
			vm.binaryOp(instruction)
		case common.OpNegate:
			vm.push(-vm.pop())
		case common.OpReturn:
			fmt.Println(vm.pop())
			return InterpretOK
		default:
			panic("unhandled default case")

		}
	}
}

func (vm *VM) push(value common.Value) {
	vm.Stack = append(vm.Stack, value)
}

func (vm *VM) pop() common.Value {
	stackTop := vm.Stack[len(vm.Stack)-1]
	vm.Stack = vm.Stack[:len(vm.Stack)-1]
	return stackTop
}

func (vm *VM) binaryOp(op common.OpCode) {
	a := vm.pop()
	b := vm.pop()

	//vm.push(a op b)
	switch op {
	case common.OpAdd:
		vm.push(a + b)
	case common.OpSubtract:
		vm.push(a - b)
	case common.OpMultiply:
		vm.push(a * b)
	case common.OpDivide:
		vm.push(a / b)
	default:
		panic("unhandled default case")
	}
}

func (vm *VM) readByte() common.OpCode {
	res := vm.Chunk.OpCodes[vm.IP]
	vm.IP++
	return res
}
