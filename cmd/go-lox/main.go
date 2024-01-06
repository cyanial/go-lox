package main

import (
	"bufio"
	"fmt"
	"github.com/cyanial/go-lox/interpreter"
	"github.com/cyanial/go-lox/vm"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		repl()
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else {
		_, _ = fmt.Fprintf(os.Stderr, "Usage: go-loc [path]\n")
		os.Exit(64)
	}
}

func repl() {
	v := vm.New()

	for {
		fmt.Print("$ ")

		sc := bufio.NewScanner(os.Stdin)
		if !sc.Scan() {
			fmt.Println()
			break
		}

		//fmt.Println("get text: ", sc.Text())
		v.Interpret(sc.Text())
	}
}

func runFile(fileName string) {
	sourceBytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	v := vm.New()

	source := string(sourceBytes)
	result := v.Interpret(source)
	if result == interpreter.CompileError {
		os.Exit(65)
	}
	if result == interpreter.RuntimeError {
		os.Exit(70)
	}
}
