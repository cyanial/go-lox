package disassembly

import (
	"fmt"
	"github.com/cyanial/go-interpreters/common"
)

func DisAssembleChunk(c *common.Chunk, name string) {
	fmt.Printf("== %s ==\n", name)

	for offset := 0; offset < len(c.OpCodes); {
		offset = DisassembleInstruction(c, offset)
	}
}

func DisassembleInstruction(c *common.Chunk, offset int) int {
	fmt.Printf("%04d ", offset)

	if offset > 0 && c.Lines[offset] == c.Lines[offset-1] {
		fmt.Printf("   | ")
	} else {
		fmt.Printf("%4d ", c.Lines[offset])
	}

	instruction := c.OpCodes[offset]
	switch instruction {
	case common.OpConstant:
		return constantInstruction("CONSTANT", c, offset)
	case common.OpConstantLong:
		return constantLongInstruction("CONSTANT_LONG", c, offset)
	case common.OpAdd:
		return simpleInstruction("ADD", offset)
	case common.OpSubtract:
		return simpleInstruction("SUBTRACT", offset)
	case common.OpMultiply:
		return simpleInstruction("MULTIPLY", offset)
	case common.OpDivide:
		return simpleInstruction("DIVIDE", offset)
	case common.OpNegate:
		return simpleInstruction("NEGATE", offset)
	case common.OpReturn:
		return simpleInstruction("RETURN", offset)
	default:
		fmt.Printf("Unknown opcode %d\n", instruction)
		return offset + 1
	}

}

func constantLongInstruction(name string, c *common.Chunk, offset int) int {
	constant := int(c.OpCodes[offset+1]) + int(c.OpCodes[offset+2])<<8 + int(c.OpCodes[offset+3])<<16
	fmt.Printf("%-16s %4d '", name, constant)
	fmt.Printf("%g", c.Constants[constant])
	fmt.Println("'")

	return offset + 4
}

func constantInstruction(name string, c *common.Chunk, offset int) int {
	constant := int(c.OpCodes[offset+1])
	fmt.Printf("%-16s %4d '", name, constant)
	fmt.Printf("%g", c.Constants[constant])
	fmt.Println("'")

	return offset + 2
}

func simpleInstruction(name string, offset int) int {
	fmt.Printf("%s\n", name)
	return offset + 1
}
