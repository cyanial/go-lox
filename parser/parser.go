package parser

import "github.com/cyanial/go-lox/token"

type Parser struct {
	Current   *token.Token
	Previous  *token.Token
	HadError  bool
	PanicMode bool
}

func New() *Parser {
	return &Parser{
		Current:   nil,
		Previous:  nil,
		HadError:  false,
		PanicMode: false,
	}
}
