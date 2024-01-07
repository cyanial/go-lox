package interpreter

import (
	"fmt"
	"github.com/cyanial/go-lox/chunk"
	"github.com/cyanial/go-lox/disassembly"
	"github.com/cyanial/go-lox/env"
	"github.com/cyanial/go-lox/op"
	"github.com/cyanial/go-lox/value"
	"os"
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
	Stack []*value.Value
}

func New(c *chunk.Chunk) *Interpreter {
	return &Interpreter{
		Chunk: c,
		IP:    0,
		Stack: make([]*value.Value, 0),
	}
}

func (it *Interpreter) Run() Result {
	for {
		if env.DebugTraceExecution {
			fmt.Printf("      STACK: ")
			for _, v := range it.Stack {
				fmt.Printf("[%s] ", v.String())
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
		case op.Nil:
			it.push(value.NewNil())
		case op.True:
			it.push(value.NewBool(true))
		case op.False:
			it.push(value.NewBool(false))
		case op.Add:
			fallthrough
		case op.Subtract:
			fallthrough
		case op.Multiply:
			fallthrough
		case op.Divide:
			if it.Stack[len(it.Stack)-1].Type != value.Number &&
				it.Stack[len(it.Stack)-2].Type != value.Number {
				it.runtimeError("operands must be numbers")
				return RuntimeError
			}
			it.binaryOp(instruction)
		case op.Not:
			v := it.pop()
			if v.Type == value.Nil {
				it.push(value.NewBool(true))
			} else if v.Type == value.Bool {
				it.push(value.NewBool(!v.AsBool()))
			} else {
				it.push(value.NewBool(v.AsNumber() == 0))
			}
		case op.Negate:
			if it.Stack[len(it.Stack)-1].Type != value.Number {
				it.runtimeError("operand must be a number")
				return RuntimeError
			}
			it.push(value.NewNumber(-it.pop().AsNumber()))
		case op.Return:
			fmt.Println(it.pop())
			return OK
		default:
			panic("unhandled default case")
		}
	}
}

func (it *Interpreter) push(value *value.Value) {
	it.Stack = append(it.Stack, value)
}

func (it *Interpreter) pop() *value.Value {
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
		it.push(value.NewNumber(a.AsNumber() + b.AsNumber()))
	case op.Subtract:
		it.push(value.NewNumber(a.AsNumber() - b.AsNumber()))
	case op.Multiply:
		it.push(value.NewNumber(a.AsNumber() * b.AsNumber()))
	case op.Divide:
		it.push(value.NewNumber(a.AsNumber() / b.AsNumber()))
	default:
		panic("unhandled default case ")
	}
}

func (it *Interpreter) runtimeError(msg string) {
	instruction := it.IP - 1
	line := it.Chunk.GetLine(instruction)
	_, _ = fmt.Fprintf(os.Stderr, "[line %d] '%s' in script\n", line, msg)
	it.resetStack()
}

func (it *Interpreter) resetStack() {
	it.Stack = make([]*value.Value, 0)
}
