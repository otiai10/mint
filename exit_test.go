// +build !freebsd

package mint_test

import (
	"os"
	"testing"

	"github.com/otiai10/mint"
)

// Exit
func TestExit(t *testing.T) {
	mint.Expect(t, func() {
		os.Exit(999999)
	}).Exit(999999)

	mint.Expect(t, func() {
		os.Exit(1)
	}).Not().Exit(0)
}