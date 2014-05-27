package mint_test

import "testing"
import "github.com/otiai10/mint"

func TestMint(t *testing.T) {
	mint.Expect(t, 1).ToBe(1)
}
func TestMint_Fail(t *testing.T) {
	// mint.Expect(t, 2).ToBe(1)
}
