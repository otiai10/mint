package mint

import "testing"
import "reflect"

type Mint struct {
	t *testing.T
}
type Testee struct {
	t        *testing.T
	actual   interface{}
	expected interface{}
	dry      bool
	not      bool
	deeply   bool
	Result   Result
}
type Result struct {
	OK      bool
	Message string
}

var (
	FailBase = 0
	FailType = 1
	Scolds   = map[int]string{
		FailBase: "Expected %sto be `%+v`, but actual `%+v`\n",
		FailType: "Expected %stype `%+v`, but actual `%T`\n",
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

// "Blend" provides (blended) *mint.Mint.
// You can save writing "t" repeatedly.
func Blend(t *testing.T) *Mint {
	return &Mint{
		t,
	}
}

// "*Mint.Expect" provides "*Testee".
// The blended mint is merely a proxy to instantiate testee.
func (m *Mint) Expect(actual interface{}) *Testee {
	return newTestee(m.t, actual)
}

// "Expect" provides "*mint.Testee".
// It has assertion methods such as "ToBe".
func Expect(t *testing.T, actual interface{}) *Testee {
	return newTestee(t, actual)
}

func newTestee(t *testing.T, actual interface{}) *Testee {
	return &Testee{t: t, actual: actual, Result: Result{OK: true}}
}
func judge(a, b interface{}, not, deeply bool) bool {
	var comparer func(a, b interface{}) bool = equal
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
