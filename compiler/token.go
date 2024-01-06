package compiler

const (
	TokenLeftParen TokenType = iota
	TokenRightParen
	TokenLeftBrace
	TokenRightBrace
	TokenComma
	TokenDot
	TokenMinus
	TokenPlus
	TokenSemicolon
	TokenSlash
	TokenStar

	TokenBang
	TokenBangEqual
	TokenEqual
	TokenEqualEqual
	TokenGreater
	TokenGreaterEqual
	TokenLess
	TokenLessEqual

	TokenIdentify
	TokenString
	TokenNumber

	TokenAnd
	TokenClass
	TokenElse
	TokenFalse
	TokenFor
	TokenFunc
	TokenIf
	TokenNil
	TokenOr
	TokenPrint
	TokenReturn
	TokenSuper
	TokenThis
	TokenTrue
	TokenVar
	TokenWhile

	TokenError
	TokenEOF
)

type TokenType uint8

type Token struct {
	Type  TokenType
	Value string
	Line  int
}
