package mint

import "testing"
import "fmt"
import "os"

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
		FailType: "Expectec %stype `%+v`, but actual `%T`\n",
	}
)

func Expect(t *testing.T, actual interface{}) *Testee {
	return &Testee{t: t, actual: actual, Result: Result{OK: true}}
}
func (testee *Testee) Dry() *Testee {
	testee.dry = true
	return testee
}
func (testee *Testee) Not() *Testee {
	testee.not = true
	return testee
}
func (testee *Testee) failed(fail ...int) *Testee {
	f := FailBase
	if 0 < len(fail) {
		f = fail[0]
	}
	return testee.failWith(f)
}
func (testee *Testee) failWith(fail int) *Testee {
	message := testee.toText(fail)
	if !testee.dry {
		fmt.Println(message)
		testee.t.Fail()
		os.Exit(1)
	}
	testee.Result.OK = false
	testee.Result.Message = message
	return testee
}
func (testee *Testee) toText(fail int) string {
	not := ""
	if testee.not {
		not = "NOT "
	}
	return fmt.Sprintf(
		Scolds[fail],
		not,
		testee.expected,
		testee.actual,
	)
}
func judge(a, b interface{}, not bool) bool {
	if not {
		return a != b
	}
	return a == b
}
