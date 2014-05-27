package mint_test

import "testing"
import "github.com/otiai10/mint"

func TestMint(t *testing.T) {
	mint.Expect(t, 1).ToBe(1)
}
func TestMint_Fail(t *testing.T) {
	// mint.Expect(t, 2).ToBe(1)
}
func TestMint_TypeOf(t *testing.T) {
	mint.Expect(t, "foo").TypeOf("string")
}
func TestMint_TypeOf_Fail(t *testing.T) {
	// mint.Expect(t, "foo").TypeOf("int")
}
