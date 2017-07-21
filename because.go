package mint

import "testing"

// Because is context printer.
func Because(t *testing.T, context string, wrapper func(*testing.T)) {
	Log("  ", context, "\n")
	wrapper(t)
}
