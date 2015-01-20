package mint

import "testing"

// Mint (mint.Mint) is wrapper for *testing.T
// blending testing type to omit repeated `t`.
type Mint struct {
	t *testing.T
}

var (
	failToBe = 0
	failType = 1
	scolds   = map[int]string{
		failToBe: "%s at line %d\nExpected %sto be\t`%+v`\nBut actual\t`%+v`",
		failType: "%s at line %d\nExpected %stype\t`%+v`\nBut actual\t`%T`",
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
	return &Testee{t: t, actual: actual, result: Result{ok: true}}
}
func judge(a, b interface{}, not, deeply bool) bool {
	comparer := getComparer(a, b, deeply)
	if not {
		return !comparer.Compare(a, b)
	}
	return comparer.Compare(a, b)
}
