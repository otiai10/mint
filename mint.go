package mint

import "testing"
import "fmt"
import "os"

type ProxyTestee struct {
	t        *testing.T
	actual   interface{}
	expected interface{}
	dry      bool
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
		FailBase: "Expected to be `%+v`, but actual `%+v`\n",
		FailType: "Expectec type `%+v`, but actual `%T`\n",
	}
)

func Expect(t *testing.T, actual interface{}) *ProxyTestee {
	return &ProxyTestee{t: t, actual: actual, Result: Result{OK: true}}
}
func (p *ProxyTestee) Dry() *ProxyTestee {
	p.dry = true
	return p
}
func (p *ProxyTestee) failed(fail ...int) *ProxyTestee {
	f := FailBase
	if 0 < len(fail) {
		f = fail[0]
	}
	return p.failWith(f)
}
func (p *ProxyTestee) failWith(fail int) *ProxyTestee {
	message := fmt.Sprintf(
		Scolds[fail],
		p.expected,
		p.actual,
	)
	if !p.dry {
		fmt.Println(message)
		p.t.Fail()
		os.Exit(1)
	}
	p.Result.OK = false
	p.Result.Message = message
	return p
}
