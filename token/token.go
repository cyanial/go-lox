package token

const (
	LeftParen Type = iota
	RightParen
	LeftBrace
	RightBrace
	Comma
	Dot
	Minus
	Plus
	Semicolon
	Slash
	Star

	Bang
	BangEqual
	Equal
	EqualEqual
	Greater
	GreaterEqual
	Less
	LessEqual

	Identify
	String
	Number

	And
	Class
	Else
	False
	For
	Func
	If
	Nil
	Or
	Print
	Return
	Super
	This
	True
	Var
	While

	Error
	EOF
)

type Type uint8

type Token struct {
	Type  Type
	Value string
	Line  int
}

func New(typ Type, value string, line int) *Token {
	return &Token{
		Type:  typ,
		Value: value,
		Line:  line,
	}
}
