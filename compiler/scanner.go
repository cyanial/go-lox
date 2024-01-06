package compiler

import (
	"unicode"
)

type Scanner struct {
	Start   int
	Current int
	Line    int
	Source  string
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		Start:   0,
		Current: 0,
		Line:    1,
		Source:  source,
	}
}

func (s *Scanner) ScanToken() *Token {
	s.skipWhitespace()
	s.Start = s.Current

	//if isAtEnd() {
	if s.Current == len(s.Source) {
		return s.makeToken(TokenEOF)
	}

	r := s.advance()
	if unicode.IsLetter(rune(r)) || r == '_' {
		return s.identifier()
	}

	if unicode.IsDigit(rune(r)) {
		return s.number()
	}

	switch r {
	case '(':
		return s.makeToken(TokenLeftParen)
	case ')':
		return s.makeToken(TokenRightParen)
	case '{':
		return s.makeToken(TokenLeftBrace)
	case '}':
		return s.makeToken(TokenRightBrace)
	case ';':
		return s.makeToken(TokenSemicolon)
	case ',':
		return s.makeToken(TokenComma)
	case '.':
		return s.makeToken(TokenDot)
	case '-':
		return s.makeToken(TokenMinus)
	case '+':
		return s.makeToken(TokenPlus)
	case '/':
		return s.makeToken(TokenSlash)
	case '*':
		return s.makeToken(TokenStar)
	case '!':
		if s.match('=') {
			return s.makeToken(TokenBangEqual)
		} else {
			return s.makeToken(TokenBang)
		}
	case '=':
		if s.match('=') {
			return s.makeToken(TokenEqualEqual)
		} else {
			return s.makeToken(TokenEqual)
		}
	case '<':
		if s.match('=') {
			return s.makeToken(TokenLessEqual)
		} else {
			return s.makeToken(TokenLess)
		}
	case '>':
		if s.match('=') {
			return s.makeToken(TokenGreaterEqual)
		} else {
			return s.makeToken(TokenGreater)
		}
	case '"':
		return s.literal()
	}

	return s.errorToken("unexpected character")
}

func (s *Scanner) makeToken(typ TokenType) *Token {
	return &Token{
		Type:  typ,
		Value: s.Source[s.Start:s.Current],
		Line:  s.Line,
	}
}

func (s *Scanner) errorToken(msg string) *Token {
	return &Token{
		Type:  TokenError,
		Value: msg,
		Line:  s.Line,
	}
}

func (s *Scanner) advance() byte {
	res := s.Source[s.Current]
	s.Current++
	return res
}

func (s *Scanner) match(expected byte) bool {
	//if isAtEnd()
	if s.Current == len(s.Source) {
		return false
	}

	if s.Source[s.Current] != expected {
		return false
	}

	s.Current++
	return true
}

func (s *Scanner) literal() *Token {
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
	return s.makeToken(TokenString)
}

func (s *Scanner) identifier() *Token {
	for s.Source[s.Current] == '_' || unicode.IsLetter(rune(s.Source[s.Current])) || unicode.IsDigit(rune(s.Source[s.Current])) {
		s.Current++
	}

	//return s.makeToken(TokenIdentify)
	return s.makeToken(s.identifierType())
}

func (s *Scanner) identifierType() TokenType {
	switch s.Source[s.Start] {
	case 'a':
		return s.checkKeyword(1, 2, "nd", TokenAnd)
	case 'c':
		return s.checkKeyword(1, 4, "lass", TokenClass)
	case 'e':
		return s.checkKeyword(1, 3, "lse", TokenElse)
	case 'i':
		return s.checkKeyword(1, 1, "f", TokenIf)
	case 'n':
		return s.checkKeyword(1, 2, "il", TokenNil)
	case 'o':
		return s.checkKeyword(1, 1, "r", TokenOr)
	case 'p':
		return s.checkKeyword(1, 4, "rint", TokenPrint)
	case 'r':
		return s.checkKeyword(1, 5, "eturn", TokenReturn)
	case 's':
		return s.checkKeyword(1, 4, "uper", TokenSuper)
	case 'v':
		return s.checkKeyword(1, 2, "ar", TokenVar)
	case 'w':
		return s.checkKeyword(1, 4, "hile", TokenWhile)
	case 'f':
		if s.Current-s.Start > 1 {
			switch s.Source[s.Start+1] {
			case 'a':
				return s.checkKeyword(2, 3, "lse", TokenFalse)
			case 'o':
				return s.checkKeyword(2, 1, "r", TokenFor)
			case 'u':
				return s.checkKeyword(2, 2, "nc", TokenFunc)
			}
		}
	case 't':
		if s.Current-s.Start > 1 {
			switch s.Source[s.Start+1] {
			case 'h':
				return s.checkKeyword(2, 2, "is", TokenThis)
			case 'r':
				return s.checkKeyword(2, 2, "ue", TokenTrue)

			}
		}
	}

	return TokenIdentify
}

func (s *Scanner) checkKeyword(start, length int, rest string, typ TokenType) TokenType {
	if s.Current-s.Start == start+length && s.Source[s.Start+start:s.Start+start+length-1] == rest {
		return typ
	}

	return TokenIdentify
}

func (s *Scanner) number() *Token {
	for unicode.IsDigit(rune(s.Source[s.Current])) {
		s.Current++
	}

	if s.Source[s.Current] == '.' && (s.Current+1 < len(s.Source) && unicode.IsDigit(rune(s.Source[s.Current+1]))) {
		// Consume the "."
		s.Current++

		for {
			for unicode.IsDigit(rune(s.Source[s.Current])) {
				s.Current++
			}
		}
	}

	return s.makeToken(TokenNumber)
}

func (s *Scanner) skipWhitespace() {
	for {
		if s.Current >= len(s.Source) {
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
