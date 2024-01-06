package disassembly

import (
	"fmt"
	"github.com/cyanial/go-lox/chunk"
	"github.com/cyanial/go-lox/op"
)

func DisassembleChunk(c *chunk.Chunk, name string) {
	fmt.Printf("***** %s *****\n", name)

	for offset := 0; offset < len(c.Codes); {
		offset = DisassembleInstruction(c, offset)
	}
}

func DisassembleInstruction(c *chunk.Chunk, offset int) int {
	fmt.Printf("%04d ", offset)

	if offset > 0 && c.GetLine(offset) == c.GetLine(offset-1) {
		fmt.Printf("     ")
	} else {
		fmt.Printf("%4d ", c.GetLine(offset))
	}

	switch instruction := c.Codes[offset]; instruction {
	case op.Constant:
		return constantInstruction("CONSTANT", c, offset)
	case op.ConstantLong:
		return constantLongInstruction("CONSTANT_LONG", c, offset)
	case op.Add:
		return simpleInstruction("ADD", offset)
	case op.Subtract:
		return simpleInstruction("SUBTRACT", offset)
	case op.Multiply:
		return simpleInstruction("MULTIPLY", offset)
	case op.Divide:
		return simpleInstruction("DIVIDE", offset)
	case op.Negate:
		return simpleInstruction("NEGATE", offset)
	case op.Return:
		return simpleInstruction("RETURN", offset)
	default:
		fmt.Printf("unknown opcode %d\n", instruction)
		return offset + 1
	}

}

func constantInstruction(name string, c *chunk.Chunk, offset int) int {
	constant := int(c.Codes[offset+1])
	fmt.Printf("%-16s %4d '", name, constant)
	fmt.Printf("%g", c.Constants[constant])
	fmt.Println("'")

	return offset + 2
}

func constantLongInstruction(name string, c *chunk.Chunk, offset int) int {
	idx := int(c.Codes[offset+1]) +
		int(c.Codes[offset+2])<<8 +
		int(c.Codes[offset+3])<<16
	fmt.Printf("%-16s %4d '", name, idx)
	fmt.Printf("%g", c.Constants[idx])
	fmt.Println("'")

	return offset + 4
}

func simpleInstruction(name string, offset int) int {
	fmt.Printf("%s\n", name)
	return offset + 1
}
