package compiler

import (
	"fmt"
	"github.com/cyanial/go-lox/chunk"
	"github.com/cyanial/go-lox/disassembly"
	"github.com/cyanial/go-lox/env"
	"github.com/cyanial/go-lox/op"
	"github.com/cyanial/go-lox/parser"
	"github.com/cyanial/go-lox/precedence"
	"github.com/cyanial/go-lox/scanner"
	"github.com/cyanial/go-lox/token"
	"github.com/cyanial/go-lox/value"
	"os"
	"strconv"
)

type Compiler struct {
	Sc *scanner.Scanner
	Pa *parser.Parser
	Ru *precedence.Rules

	Ck *chunk.Chunk
}

func New() *Compiler {
	return &Compiler{}
}

func (c *Compiler) Compile(source string) (*chunk.Chunk, bool) {
	c.Sc = scanner.New(source)
	c.Pa = parser.New()
	c.Ru = precedence.NewRules(c.unary, c.binary, c.grouping, c.number, c.literal)
	c.Ck = chunk.New()

	c.advance()
	c.expression()
	c.consume(token.EOF, "expect end of expression")
	c.endCompiler()

	return c.Ck, true
}

func (c *Compiler) advance() {
	c.Pa.Previous = c.Pa.Current

	for {
		c.Pa.Current = c.Sc.ScanToken()
		if c.Pa.Current.Type != token.Error {
			break
		}

		c.errorAtCurrent(c.Pa.Current.Value)
	}
}

func (c *Compiler) expression() {
	c.parsePrecedence(precedence.Assignment)
}

func (c *Compiler) consume(typ token.Type, msg string) {
	if c.Pa.Current.Type == typ {
		c.advance()
		return
	}

	c.errorAtCurrent(msg)
}

func (c *Compiler) endCompiler() {
	if env.DebugPrintCode && !c.Pa.HadError {
		disassembly.DisassembleChunk(c.Ck, "code")
	}
}

func (c *Compiler) parsePrecedence(p precedence.Precedence) {
	c.advance()
	prefixRule := c.Ru.Get(c.Pa.Previous.Type).Prefix
	if prefixRule == nil {
		c.error("expect expression")
		return
	}

	prefixRule()

	for p <= c.Ru.Get(c.Pa.Current.Type).Precedence {
		c.advance()
		infixRule := c.Ru.Get(c.Pa.Previous.Type).Infix
		infixRule()
	}
}

func (c *Compiler) grouping() {
	c.expression()
	c.consume(token.RightParen, "expect ')' after expression")
}

func (c *Compiler) literal() {
	switch c.Pa.Previous.Type {
	case token.False:
		c.emitByte(op.False)
	case token.Nil:
		c.emitByte(op.Nil)
	case token.True:
		c.emitByte(op.True)
	default:
		return
	}
}

func (c *Compiler) number() {
	valueFloat64, _ := strconv.ParseFloat(c.Pa.Previous.Value, 64)
	c.emitConstant(value.NewNumber(valueFloat64))
}

func (c *Compiler) unary() {
	typ := c.Pa.Previous.Type

	//c.expression()
	c.parsePrecedence(precedence.Unary)

	switch typ {
	case token.Minus:
		c.emitByte(op.Negate)
	default:
		return
	}
}

func (c *Compiler) binary() {
	typ := c.Pa.Previous.Type
	rule := c.Ru.Get(typ)

	c.parsePrecedence(rule.Precedence + 1)

	switch typ {
	case token.Plus:
		c.emitByte(op.Add)
	case token.Minus:
		c.emitByte(op.Subtract)
	case token.Star:
		c.emitByte(op.Multiply)
	case token.Slash:
		c.emitByte(op.Divide)
	default:
		return
	}

}

func (c *Compiler) emitConstant(v *value.Value) {
	c.Ck.AddConstant(v, c.Pa.Previous.Line)
}

func (c *Compiler) emitBytes(ops ...op.Code) {
	for _, o := range ops {
		c.emitByte(o)
	}
}

func (c *Compiler) emitByte(o op.Code) {
	c.Ck.AddOp(o, c.Pa.Previous.Line)
}

func (c *Compiler) emitReturn() {
	c.emitByte(op.Return)
}

func (c *Compiler) errorAtCurrent(msg string) {
	c.errorAt(c.Pa.Current, msg)
}

func (c *Compiler) error(msg string) {
	c.errorAt(c.Pa.Previous, msg)
}

func (c *Compiler) errorAt(tk *token.Token, msg string) {
	if c.Pa.PanicMode {
		return
	}
	c.Pa.PanicMode = true
	_, _ = fmt.Fprintf(os.Stderr, "[line %d] error", tk.Line)

	if tk.Type == token.EOF {
		_, _ = fmt.Fprint(os.Stderr, " at end")
	} else if tk.Type == token.Error {
		// nothing
	} else {
		_, _ = fmt.Fprintf(os.Stderr, " at '%s'", tk.Value)
	}

	_, _ = fmt.Fprintf(os.Stderr, ": %s\n", msg)
	c.Pa.HadError = true
}
