package mint_test

import "github.com/otiai10/mint"
import "testing"

// "Expect" provides "*mint.Testee".
// It has assertion methods such as "ToBe".
func ExampleExpect(t *testing.T) {
	mint.Expect(t, 100).ToBe(100)
	mint.Expect(t, 100).TypeOf("int")
}

// "*Testee.ToBe" can assert the testee to equal the parameter of this func.
// OS will exit with code 1, when the assertion fail.
// If you don't want to exit, see "Dry()".
func ExampleTestee_ToBe(t *testing.T) {
	mint.Expect(t, 100).ToBe(100)
}

// "*Testee.TypeOf" can assert the type of testee to equal the parameter of this func.
// OS will exit with code 1, when the assertion fail.
// If you don't want to exit, see "Dry()".
func ExampleTestee_TypeOf(t *testing.T) {
	mint.Expect(t, 100).TypeOf("int")
}

// "*Testee.Not" makes following assertion conversed.
func ExampleTestee_Not(t *testing.T) {
	mint.Expect(t, 100).Not().ToBe(200)
	mint.Expect(t, 100).Not().TypeOf("string")
}

// "*Testee.Dry" makes the testee NOT to call "os.Exit(1)".
// Use this if you want to fail test in a purpose.
func ExampleTestee_Dry(t *testing.T) {
	result := mint.Expect(t, 100).Dry().ToBe(100).Result
	if !result.OK {
		t.Fail()
	}
}

// "Blend" provides (blended) *mint.Mint.
// You can save writing "t" repeatedly.
func ExampleBlend(t *testing.T) {
	// get blended mint
	m := mint.Blend(t)

	m.Expect(100).ToBe(100)
	m.Expect(100).Not().ToBe(200)
}
