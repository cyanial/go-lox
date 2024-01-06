package interpreter

import (
	"fmt"
	"github.com/cyanial/go-lox/chunk"
	"github.com/cyanial/go-lox/disassembly"
	"github.com/cyanial/go-lox/op"
	"github.com/cyanial/go-lox/value"
	"os"
)

var (
	debugPrintCode      = false
	debugTraceExecution = false
)

const (
	OK Result = iota
	CompileError
	RuntimeError
)

type Result int

type Interpreter struct {
	Chunk *chunk.Chunk
	IP    int
	Stack []value.Value
}

func init() {
	if dpc := os.Getenv("debugPrintCode"); dpc == "true" {
		debugPrintCode = true
	}
	if dte := os.Getenv("debugTraceExecution"); dte == "true" {
		debugTraceExecution = true
	}
}

func New(c *chunk.Chunk) *Interpreter {
	return &Interpreter{
		Chunk: c,
		IP:    0,
		Stack: make([]value.Value, 0),
	}
}

func (it *Interpreter) Run() Result {
	for {
		if debugTraceExecution {
			fmt.Printf("      STACK: ")
			for _, slot := range it.Stack {
				fmt.Printf("[%v] ", slot)
			}
			fmt.Println()
			disassembly.DisassembleInstruction(it.Chunk, it.IP)
		}

		switch instruction := it.readByte(); instruction {
		case op.Constant:
			it.push(it.Chunk.Constants[it.readByte()])
		case op.ConstantLong:
			idx := 0
			idx += int(it.readByte())
			idx += int(it.readByte()) << 8
			idx += int(it.readByte()) << 16
			it.push(it.Chunk.Constants[idx])
		case op.Add:
			fallthrough
		case op.Subtract:
			fallthrough
		case op.Multiply:
			fallthrough
		case op.Divide:
			it.binaryOp(instruction)
		case op.Negate:
			it.push(-it.pop())
		case op.Return:
			fmt.Println(it.pop())
			return OK
		default:
			panic("unhandled default case")
		}
	}
}

func (it *Interpreter) push(value value.Value) {
	it.Stack = append(it.Stack, value)
}

func (it *Interpreter) pop() value.Value {
	top := it.Stack[len(it.Stack)-1]
	it.Stack = it.Stack[:len(it.Stack)-1]
	return top
}

func (it *Interpreter) readByte() op.Code {
	o := it.Chunk.Codes[it.IP]
	it.IP++
	return o
}

func (it *Interpreter) binaryOp(o op.Code) {
	a := it.pop()
	b := it.pop()

	// it.push(a o b)
	switch o {
	case op.Add:
		it.push(a + b)
	case op.Subtract:
		it.push(a - b)
	case op.Multiply:
		it.push(a * b)
	case op.Divide:
		it.push(a / b)
	default:
		panic("unhandled default case ")
	}
}
