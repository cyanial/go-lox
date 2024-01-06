package compiler

import "fmt"

func Compile(source string) {
	//scanner := NewScanner(source)

	line := -1

	for {
		//token := scanToken()
		var token Token
		if token.Line != line {
			fmt.Printf("%4d ", token.Line)
			line = token.Line
		} else {
			fmt.Printf("   | ")
		}
		//fmt.Printf("%2d '%.*s'\n", token.Type, token.Length, token.Start)
		fmt.Printf("%2d '%s'\n", token.Type, token.Value)

		if token.Type == TokenEOF {
			break
		}
	}
}
