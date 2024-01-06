package main

import (
	"bufio"
	"fmt"
	"github.com/cyanial/go-interpreters/vm"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		repl()
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else {
		_, _ = fmt.Fprintf(os.Stderr, "Usage: ./goi [path]\n")
		os.Exit(64)
	}
}

func repl() {

	scanner := bufio.NewScanner(os.Stdin)
	//v := vm.New(common.NewChunk(), false)

	for {
		fmt.Print("> ")
		for scanner.Scan() {
			line := scanner.Text()

			_ = line
			//interpret(line)
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}
}

func runFile(fileName string) {
	sourceBytes, err := os.ReadFile(fileName)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "could not open file %s.\n", fileName)
		os.Exit(74)
	}

	source := string(sourceBytes)
	_ = source
	//result := interpret(source)
	var result vm.InterpretResult

	if result == vm.InterpretCompileError {
		os.Exit(65)
	}

	if result == vm.InterpretRuntimeError {
		os.Exit(70)
	}
}
