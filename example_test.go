package mint_test

import "github.com/otiai10/mint"
import "testing"

func ExampleExpect(t *testing.T) {
	mint.Expect(t, 100).ToBe(100)
	mint.Expect(t, 100).TypeOf("int")
}

func ExampleTestee_ToBe(t *testing.T) {
	mint.Expect(t, 100).ToBe(100)
}

func ExampleTestee_TypeOf(t *testing.T) {
	mint.Expect(t, 100).TypeOf("int")
}

func ExampleTestee_Not(t *testing.T) {
	mint.Expect(t, 100).Not().ToBe(200)
	mint.Expect(t, 100).Not().TypeOf("string")
}

func ExampleTestee_Dry(t *testing.T) {
	result := mint.Expect(t, 100).Dry().ToBe(100).Result
	if !result.OK {
		t.Fail()
	}
}

func ExampleBlend(t *testing.T) {
	// get blended mint
	m := mint.Blend(t)

	m.Expect(100).ToBe(100)
	m.Expect(100).Not().ToBe(200)
}
