package scanner

import (
	"github.com/cyanial/go-lox/token"
	"unicode"
)

type Scanner struct {
	Start   int
	Current int
	Line    int
	Source  string
}

func New(source string) *Scanner {
	return &Scanner{
		Start:   0,
		Current: 0,
		Line:    1,
		Source:  source,
	}
}

func (s *Scanner) ScanToken() *token.Token {

	s.skipWhitespace()
	s.Start = s.Current

	if s.Current == len(s.Source) {
		return s.makeToken(token.EOF)
	}

	c := s.advance()

	if unicode.IsLetter(rune(c)) || c == '_' {
		return s.identifier()
	}

	if unicode.IsDigit(rune(c)) {
		return s.number()
	}

	switch c {
	case '(':
		return s.makeToken(token.LeftParen)
	case ')':
		return s.makeToken(token.RightParen)
	case '{':
		return s.makeToken(token.LeftBrace)
	case '}':
		return s.makeToken(token.RightBrace)
	case ';':
		return s.makeToken(token.Semicolon)
	case ',':
		return s.makeToken(token.Comma)
	case '.':
		return s.makeToken(token.Dot)
	case '-':
		return s.makeToken(token.Minus)
	case '+':
		return s.makeToken(token.Plus)
	case '/':
		return s.makeToken(token.Slash)
	case '*':
		return s.makeToken(token.Star)
	case '!':
		if s.match('=') {
			return s.makeToken(token.BangEqual)
		}
		return s.makeToken(token.Bang)
	case '=':
		if s.match('=') {
			return s.makeToken(token.EqualEqual)
		}
		return s.makeToken(token.Equal)
	case '<':
		if s.match('=') {
			return s.makeToken(token.LessEqual)
		}
		return s.makeToken(token.Less)
	case '>':
		if s.match('=') {
			return s.makeToken(token.GreaterEqual)
		}
		return s.makeToken(token.Greater)
	case '"':
		return s.literal()
	}

	return s.errorToken("unexpected character")
}

func (s *Scanner) advance() byte {
	r := s.Source[s.Current]
	s.Current++
	return r
}

func (s *Scanner) match(expected byte) bool {
	if s.Current == len(s.Source) {
		return false
	}

	if s.Source[s.Current] != expected {
		return false
	}

	s.Current++
	return true
}

func (s *Scanner) skipWhitespace() {

	for {
		if s.Current == len(s.Source) {
			return
		}

		switch c := s.Source[s.Current]; c {
		case ' ':
			fallthrough
		case '\r':
			fallthrough
		case '\t':
			s.Current++
		case '\n':
			s.Line++
			s.Current++
		case '/':
			if s.Current+1 < len(s.Source) && s.Source[s.Current+1] == '/' {
				// a comment goes util the end of the line
				for s.Current < len(s.Source) && s.Source[s.Current] != '\n' {
					s.Current++
				}
			} else {
				return
			}
		default:
			return
		}
	}
}

func (s *Scanner) makeToken(typ token.Type) *token.Token {
	return token.New(typ, s.Source[s.Start:s.Current], s.Line)
}

func (s *Scanner) errorToken(msg string) *token.Token {
	return token.New(token.Error, msg, s.Line)
}

func (s *Scanner) literal() *token.Token {
	for s.Current < len(s.Source) && s.Source[s.Current] != '"' {
		if s.Source[s.Current] == '\n' {
			s.Line++
		}
		s.Current++
	}

	if s.Current == len(s.Source) {
		return s.errorToken("unterminated string")
	}

	// closing quote
	s.Current++
	return s.makeToken(token.String)
}

func (s *Scanner) number() *token.Token {
	for unicode.IsDigit(rune(s.Source[s.Current])) {
		s.Current++
		if s.Source[s.Current] == '.' &&
			(s.Current+1 < len(s.Source) && unicode.IsDigit(rune(s.Source[s.Current+1]))) {
			// Consume the "."
			s.Current++

			for {
				for unicode.IsDigit(rune(s.Source[s.Current])) {
					s.Current++
				}
			}
		}
	}

	return s.makeToken(token.Number)
}

func (s *Scanner) identifier() *token.Token {
	for s.Source[s.Current] == '_' ||
		unicode.IsLetter(rune(s.Source[s.Current])) ||
		unicode.IsDigit(rune(s.Source[s.Current])) {
		s.Current++
	}

	//return s.makeToken(TokenIdentify)
	return s.makeToken(s.identifierType())
}

func (s *Scanner) identifierType() token.Type {

	switch s.Source[s.Start] {
	case 'a':
		return s.checkKeyword(1, 2, "nd", token.And)
	case 'c':
		return s.checkKeyword(1, 4, "lass", token.Class)
	case 'e':
		return s.checkKeyword(1, 3, "lse", token.Else)
	case 'i':
		return s.checkKeyword(1, 1, "f", token.If)
	case 'n':
		return s.checkKeyword(1, 2, "il", token.Nil)
	case 'o':
		return s.checkKeyword(1, 1, "r", token.Or)
	case 'p':
		return s.checkKeyword(1, 4, "rint", token.Print)
	case 'r':
		return s.checkKeyword(1, 5, "eturn", token.Return)
	case 's':
		return s.checkKeyword(1, 4, "uper", token.Super)
	case 'v':
		return s.checkKeyword(1, 2, "ar", token.Var)
	case 'w':
		return s.checkKeyword(1, 4, "hile", token.While)
	case 'f':
		if s.Start+1 < len(s.Source) {
			switch s.Source[s.Start+1] {
			case 'a':
				return s.checkKeyword(2, 3, "lse", token.False)
			case 'o':
				return s.checkKeyword(2, 1, "r", token.For)
			case 'u':
				return s.checkKeyword(2, 2, "nc", token.Func)
			}
		}
	case 't':
		if s.Start+1 < len(s.Source) {
			switch s.Source[s.Start+1] {
			case 'h':
				return s.checkKeyword(2, 2, "is", token.This)
			case 'r':
				return s.checkKeyword(2, 2, "ue", token.True)

			}
		}
	}

	return token.Identify
}

func (s *Scanner) checkKeyword(start, length int, rest string, typ token.Type) token.Type {
	if s.Start+start+length-1 < len(s.Source) && s.Source[s.Start+start:s.Start+start+length-1] == rest {
		return typ
	}

	return token.Identify
}
