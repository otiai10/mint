package mint

import (
	"fmt"
	"testing"
)

// Because is context printer.
func Because(t *testing.T, context string, wrapper func(*testing.T)) {
	fmt.Print(context + "\n")
	wrapper(t)
}
