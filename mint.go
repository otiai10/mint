package mint

import "testing"

type Mint struct {
	t *testing.T
}
type Testee struct {
	t        *testing.T
	actual   interface{}
	expected interface{}
	dry      bool
	not      bool
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

func Blend(t *testing.T) *Mint {
	return &Mint{
		t,
	}
}
func Expect(t *testing.T, actual interface{}) *Testee {
	return newTestee(t, actual)
}
func (m *Mint) Expect(actual interface{}) *Testee {
	return newTestee(m.t, actual)
}
func newTestee(t *testing.T, actual interface{}) *Testee {
	return &Testee{t: t, actual: actual, Result: Result{OK: true}}
}
func judge(a, b interface{}, not bool) bool {
	if not {
		return a != b
	}
	return a == b
}
