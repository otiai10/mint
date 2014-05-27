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
