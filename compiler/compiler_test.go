package compiler

import (
	"github.com/cyanial/go-lox/env"
	"testing"
)

func TestCompilerExpression(t *testing.T) {
	c := New()

	tests := []struct {
		source string
	}{
		{
			source: "1 + 1",
		},
		{
			source: "1 - 1",
		},
		{
			source: "1 / 1",
		},
		{
			source: "1 * 1",
		},
		{
			source: "(2+3)*5",
		},
		{
			source: "2.2+3.1*51.2/(2.2-1.0)",
		},
	}

	env.DebugPrintCode = true
	env.DebugTraceExecution = true

	for _, tt := range tests {
		_, _ = c.Compile(tt.source)
	}
}
