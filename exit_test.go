// +build !freebsd

package mint_test

import (
	"log"
	"os"
	"testing"

	"github.com/otiai10/mint"
)

// Exit
func TestExit(t *testing.T) {
	mint.Expect(t, func() {
		log.Fatalln("intentionally failed")
	}).Exit(1)

	mint.Expect(t, func() {
		os.Exit(1)
	}).Not().Exit(0)
}
