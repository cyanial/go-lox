package token

import "fmt"

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

func (t *Token) String() string {
	res := "["
	switch t.Type {
	case LeftParen:
		res += "LeftParen"
	case RightParen:
		res += "RightParen"
	case LeftBrace:
		res += "LeftBrace"
	case RightBrace:
		res += "RightBrace"
	case Comma:
		res += "Comma"
	case Dot:
		res += "Dot"
	case Minus:
		res += "Minus"
	case Plus:
		res += "Plus"
	case Semicolon:
		res += "Semicolon"
	case Slash:
		res += "Slash"
	case Star:
		res += "Star"
	case Bang:
		res += "Bang"
	case BangEqual:
		res += "BangEqual"
	case Equal:
		res += "Equal"
	case EqualEqual:
		res += "EqualEqual"
	case Greater:
		res += "Greater"
	case GreaterEqual:
		res += "GreaterEqual"
	case Less:
		res += "Less"
	case LessEqual:
		res += "LessEqual"
	case Identify:
		res += "Identify"
	case String:
		res += "String"
	case Number:
		res += "Number"
	case And:
		res += "And"
	case Class:
		res += "Class"
	case Else:
		res += "Else"
	case False:
		res += "False"
	case For:
		res += "For"
	case Func:
		res += "Func"
	case If:
		res += "If"
	case Nil:
		res += "Nil"
	case Or:
		res += "Or"
	case Print:
		res += "Print"
	case Return:
		res += "Return"
	case Super:
		res += "Super"
	case This:
		res += "This"
	case True:
		res += "True"
	case Var:
		res += "Var"
	case While:
		res += "While"
	case Error:
		res += "Error"
	case EOF:
		res += "EOF"
	}
	res += ", "
	res += t.Value
	res += fmt.Sprintf(", %d]", t.Line)
	return res
}
