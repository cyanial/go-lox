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
		{
			source: "(-1 + 2) * 3 - -4",
		},
	}

	env.DebugPrintCode = true
	env.DebugTraceExecution = true

	for _, tt := range tests {
		_, _ = c.Compile(tt.source)
	}
}

func TestCompilerBoolNil(t *testing.T) {
	c := New()

	tests := []struct {
		source string
	}{
		{
			source: "true",
		},
		{
			source: "false",
		},
		{
			source: "nil",
		},
	}

	env.DebugPrintCode = true
	env.DebugTraceExecution = true

	for _, tt := range tests {
		_, _ = c.Compile(tt.source)
	}
}

func TestCompilerLogicalNotAndFalseValue(t *testing.T) {
	c := New()

	tests := []struct {
		source string
	}{
		{
			source: "!nil",
		},
		{
			source: "!true",
		},
		{
			source: "!false",
		},
		{
			source: "!0",
		},
		{
			source: "!123.2",
		},
	}

	env.DebugPrintCode = true
	env.DebugTraceExecution = true

	for _, tt := range tests {
		_, _ = c.Compile(tt.source)
	}
}

func TestCompilerCompareEquality(t *testing.T) {
	c := New()

	tests := []struct {
		source string
	}{
		{
			source: "!(5 - 4 > 3 * 2 == !nil)",
		},
	}

	env.DebugPrintCode = true
	env.DebugTraceExecution = true

	for _, tt := range tests {
		_, _ = c.Compile(tt.source)
	}
}
