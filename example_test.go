package mint_test

import "github.com/otiai10/mint"
import "testing"

// This is example of `Expect` and `ToBe`
// `ToBe` can assert to eqaal
func Example_Expect_ToBe(t *testing.T) {
	yours := 100
	expected := 100
	mint.Expect(t, yours).ToBe(expected)
}

// This is example of `Dry`
// `Dry` does NOT os.Exit(1)
// but returns `ProxyTestee` with `Result`
func Example_Dry(t *testing.T) {
	yours := 100
	expected := 100
	result := mint.Expect(t, yours).Dry().ToBe(expected).Result
	if !result.OK {
		t.Fail()
	}
}

// `Blend` provide *mint.Testee.
// It can save the repeating of `t`.
func ExampleMint_Expect(t *testing.T) {
	// get blended mint
	m := mint.Blend(t)

	yours := 100
	m.Expect(yours).ToBe(100)
	result := m.Expect(yours).Dry().ToBe(200).Result
	if result.OK {
		t.Fail()
	}
	m.Expect(yours).Not().ToBe(200)
}
