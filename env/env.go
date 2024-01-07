package env

import "os"

var (
	DebugPrintCode      = false
	DebugTraceExecution = false
)

func init() {
	if dpc := os.Getenv("debugPrintCode"); dpc == "true" {
		DebugPrintCode = true
	}
	if dte := os.Getenv("debugTraceExecution"); dte == "true" {
		DebugTraceExecution = true
	}
}
