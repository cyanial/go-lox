package compiler

//
//import (
//	"fmt"
//	"github.com/cyanial/go-interpreters/common"
//	"github.com/cyanial/go-interpreters/disassembly"
//	"math"
//	"os"
//	"strconv"
//)
//
//type Compiler struct {
//	Chunk   *common.Chunk
//	Parser  *Parser
//	Scanner *Scanner
//	Rules   map[TokenType]*ParserRule
//
//	debugPrintCode bool
//}
//
//func NewCompiler(source string, debugPrintCode bool) *Compiler {
//	c := &Compiler{
//		Scanner:        NewScanner(source),
//		Parser:         &Parser{},
//		debugPrintCode: debugPrintCode,
//	}
//
//	c.Rules = map[TokenType]*ParserRule{
//		TokenLeftParen:    {c.grouping, nil, PrecedenceNone},
//		TokenRightParen:   {nil, nil, PrecedenceNone},
//		TokenLeftBrace:    {nil, nil, PrecedenceNone},
//		TokenRightBrace:   {nil, nil, PrecedenceNone},
//		TokenComma:        {nil, nil, PrecedenceNone},
//		TokenDot:          {nil, nil, PrecedenceNone},
//		TokenMinus:        {c.unary, c.binary, PrecedenceTerm},
//		TokenPlus:         {nil, c.binary, PrecedenceTerm},
//		TokenSemicolon:    {nil, nil, PrecedenceNone},
//		TokenSlash:        {nil, c.binary, PrecedenceFactor},
//		TokenStar:         {nil, c.binary, PrecedenceFactor},
//		TokenBang:         {nil, nil, PrecedenceNone},
//		TokenBangEqual:    {nil, nil, PrecedenceNone},
//		TokenEqual:        {nil, nil, PrecedenceNone},
//		TokenEqualEqual:   {nil, nil, PrecedenceNone},
//		TokenGreater:      {nil, nil, PrecedenceNone},
//		TokenGreaterEqual: {nil, nil, PrecedenceNone},
//		TokenLess:         {nil, nil, PrecedenceNone},
//		TokenLessEqual:    {nil, nil, PrecedenceNone},
//		TokenIdentify:     {nil, nil, PrecedenceNone},
//		TokenString:       {nil, nil, PrecedenceNone},
//		TokenNumber:       {c.number, nil, PrecedenceNone},
//		TokenAnd:          {nil, nil, PrecedenceNone},
//		TokenClass:        {nil, nil, PrecedenceNone},
//		TokenElse:         {nil, nil, PrecedenceNone},
//		TokenFalse:        {nil, nil, PrecedenceNone},
//		TokenFor:          {nil, nil, PrecedenceNone},
//		TokenFunc:         {nil, nil, PrecedenceNone},
//		TokenIf:           {nil, nil, PrecedenceNone},
//		TokenNil:          {nil, nil, PrecedenceNone},
//		TokenOr:           {nil, nil, PrecedenceNone},
//		TokenPrint:        {nil, nil, PrecedenceNone},
//		TokenReturn:       {nil, nil, PrecedenceNone},
//		TokenSuper:        {nil, nil, PrecedenceNone},
//		TokenThis:         {nil, nil, PrecedenceNone},
//		TokenTrue:         {nil, nil, PrecedenceNone},
//		TokenVar:          {nil, nil, PrecedenceNone},
//		TokenWhile:        {nil, nil, PrecedenceNone},
//		TokenError:        {nil, nil, PrecedenceNone},
//		TokenEOF:          {nil, nil, PrecedenceNone},
//	}
//
//	return c
//}
//
//func (c *Compiler) Compile() bool {
//	c.advance()
//	c.expression()
//	c.consume(TokenEOF, "expect end of expression")
//	c.endCompiler()
//	return !c.Parser.HadError
//}
//
//func (c *Compiler) advance() {
//	c.Parser.Previous = c.Parser.Current
//
//	for {
//		c.Parser.Current = c.Scanner.ScanToken()
//		if c.Parser.Current.Type != TokenError {
//			break
//		}
//
//		c.errorAtCurrent("")
//	}
//}
//
//func (c *Compiler) expression() {
//	c.parsePrecedence(PrecedenceAssignment)
//}
//
//func (c *Compiler) consume(typ TokenType, msg string) {
//	if c.Parser.Current.Type == typ {
//		c.advance()
//		return
//	}
//
//	c.errorAtCurrent(msg)
//}
//
//func (c *Compiler) endCompiler() {
//	c.emitByte(common.OpReturn)
//
//	if c.debugPrintCode && !c.Parser.HadError {
//		disassembly.DisAssembleChunk(c.Chunk, "code")
//	}
//}
//
//func (c *Compiler) grouping() {
//	c.expression()
//	c.consume(TokenRightParen, "Expect ')' after expression")
//}
//
//func (c *Compiler) number() {
//	value, _ := strconv.ParseFloat(c.Parser.Previous.Value, 64)
//	c.emitConstant(common.Value(value))
//}
//
//func (c *Compiler) unary() {
//	opTyp := c.Parser.Previous.Type
//
//	c.parsePrecedence(PrecedenceUnary)
//
//	switch opTyp {
//	case TokenMinus:
//		c.emitByte(common.OpNegate)
//	default:
//		return
//	}
//}
//
//func (c *Compiler) binary() {
//	opTyp := c.Parser.Previous.Type
//	rule := c.getRule(opTyp)
//
//	c.parsePrecedence(rule.Precedence + 1)
//
//	switch opTyp {
//	case TokenPlus:
//		c.emitByte(common.OpAdd)
//	case TokenMinus:
//		c.emitByte(common.OpSubtract)
//	case TokenStar:
//		c.emitByte(common.OpMultiply)
//	case TokenSlash:
//		c.emitByte(common.OpDivide)
//	default:
//		return
//	}
//}
//
//func (c *Compiler) parsePrecedence(precedence Precedence) {
//	c.advance()
//	prefixRule := c.getRule(c.Parser.Previous.Type).Prefix
//	if prefixRule == nil {
//		//error("expect expression")
//		_, _ = fmt.Fprintf(os.Stderr, "expect expression")
//		return
//	}
//
//	prefixRule()
//
//	for precedence <= c.getRule(c.Parser.Current.Type).Precedence {
//		c.advance()
//		infixRule := c.getRule(c.Parser.Previous.Type).infix
//		infixRule()
//	}
//}
//
//func (c *Compiler) getRule(typ TokenType) *ParserRule {
//	return c.Rules[typ]
//}
//
//func (c *Compiler) emitConstant(value common.Value) {
//	c.emitBytes(common.OpConstant, c.makeConstant(value))
//}
//
//func (c *Compiler) makeConstant(value common.Value) common.OpCode {
//	constant := c.Chunk.AddConstant(value)
//	if constant > math.MaxUint8 {
//		_, _ = fmt.Fprint(os.Stderr, "too many constants in one chunk")
//		return 0
//	}
//
//	return common.OpCode(constant)
//}
//
//func (c *Compiler) emitByte(b common.OpCode) {
//	c.Chunk.AddOp(b, c.Parser.Previous.Line)
//}
//
//func (c *Compiler) emitBytes(bs ...common.OpCode) {
//	for _, b := range bs {
//		c.emitByte(b)
//	}
//}
//
//func (c *Compiler) errorAtCurrent(msg string) {
//	c.errorAt(c.Parser.Current, msg)
//}
//
//func (c *Compiler) errorAt(token *Token, msg string) {
//	if c.Parser.PanicMode {
//		return
//	}
//	c.Parser.PanicMode = true
//
//	_, _ = fmt.Fprintf(os.Stderr, "[line %d] Error", token.Line)
//
//	if token.Type == TokenEOF {
//		_, _ = fmt.Fprint(os.Stderr, " at end")
//	} else if token.Type == TokenError {
//		// nothing
//	} else {
//		_, _ = fmt.Fprintf(os.Stderr, " at '%s'", token.Value)
//	}
//
//	_, _ = fmt.Fprintf(os.Stderr, ": %s\n", msg)
//	c.Parser.HadError = true
//}
