package mint

import "testing"
import "reflect"

// Mint (mint.Mint) is wrapper for *testing.T
// blending testing type to omit repeated `t`.
type Mint struct {
	t *testing.T
}

// Testee is holder of interfaces which user want to assert
// and also has its result.
type Testee struct {
	t        *testing.T
	actual   interface{}
	expected interface{}
	dry      bool
	not      bool
	deeply   bool
	Result   Result
}

// Result provide the results of assertion
// for `Dry` option.
type Result struct {
	OK      bool
	Message string
}

var (
	failToBe = 0
	failType = 1
	scolds   = map[int]string{
		failToBe: "Expected %sto be `%+v`, but actual `%+v`\n",
		failType: "Expected %stype `%+v`, but actual `%T`\n",
	}
)
var (
	redB     = "\033[1;31m"
	reset    = "\033[0m"
	colorize = map[string]func(string) string{
		"red": func(v string) string {
			return redB + v + reset
		},
	}
)

// Blend provides (blended) *mint.Mint.
// You can save writing "t" repeatedly.
func Blend(t *testing.T) *Mint {
	return &Mint{
		t,
	}
}

// Expect provides "*Testee".
// The blended mint is merely a proxy to instantiate testee.
func (m *Mint) Expect(actual interface{}) *Testee {
	return newTestee(m.t, actual)
}

// Expect provides "*mint.Testee".
// It has assertion methods such as "ToBe".
func Expect(t *testing.T, actual interface{}) *Testee {
	return newTestee(t, actual)
}

func newTestee(t *testing.T, actual interface{}) *Testee {
	return &Testee{t: t, actual: actual, Result: Result{OK: true}}
}
func judge(a, b interface{}, not, deeply bool) bool {
	comparer := equal
	if deeply {
		comparer = deepEqual
	}
	if not {
		return !comparer(a, b)
	}
	return comparer(a, b)
}
func equal(a, b interface{}) bool {
	return a == b
}
func deepEqual(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}
