package mint_test

import "testing"
import "github.com/otiai10/mint"

func TestMint_ToBe(t *testing.T) {
	mint.Expect(t, 1).ToBe(1)
}
func TestMint_ToBe_Fail(t *testing.T) {
	r := mint.Expect(t, 2).Dry().ToBe(1).Result
	// assert mint by using mint
	mint.Expect(t, r.OK).ToBe(false)
}

type MyStruct struct{}

func TestMint_TypeOf(t *testing.T) {
	mint.Expect(t, "foo").TypeOf("string")

	bar := MyStruct{}
	mint.Expect(t, bar).TypeOf("mint_test.MyStruct")
}
func TestMint_TypeOf_Fail(t *testing.T) {
	r := mint.Expect(t, "foo").Dry().TypeOf("int").Result
	// assert mint by using mint
	mint.Expect(t, r.OK).ToBe(false)

	bar := MyStruct{}
	r = mint.Expect(t, bar).Dry().TypeOf("foo.Bar").Result
	// assert mint by using mint
	mint.Expect(t, r.OK).ToBe(false)
	mint.Expect(t, r.Message).ToBe("Expected type `foo.Bar`, but actual `mint_test.MyStruct`\n")
}

func TestMint_Not(t *testing.T) {
	mint.Expect(t, 100).Not().ToBe(200)
	mint.Expect(t, "foo").Not().TypeOf("int")
	mint.Expect(t, true).Not().ToBe(nil)
}
func TestMint_Not_Fail(t *testing.T) {
	r := mint.Expect(t, "foo").Dry().Not().TypeOf("string").Result
	// assert mint by using mint
	mint.Expect(t, r.OK).Not().ToBe(true)
}

func TestMint_Deeply(t *testing.T) {
	map0 := &map[int]string{
		3:  "three",
		5:  "five",
		10: "ten",
	}
	map1 := &map[int]string{
		3:  "three",
		5:  "five",
		10: "ten",
	}
	// It SHALLOWLY different.
	mint.Expect(t, map0).Not().ToBe(map1)
	// But it DEEPLY equal.
	mint.Expect(t, map0).Deeply().ToBe(map1)
}

// Blend is a shorhand to get testee
func TestMint_Blend(t *testing.T) {
	m := mint.Blend(t)
	// assert mint by using mint
	mint.Expect(t, m).TypeOf("*mint.Mint")
	mint.Expect(t, m.Expect("foo")).TypeOf("*mint.Testee")
}
