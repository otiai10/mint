package mint

import "testing"
import "fmt"
import "os"

type ProxyTestee struct {
	t        *testing.T
	actual   interface{}
	expected interface{}
}

var (
	FailBase = 0
	FailType = 1
)
var Scolds = map[int]string{
	FailBase: "Expected to be `%+v`, but actual `%+v`\n",
	FailType: "Expectec type `%+v`, but actual `%T`\n",
}

func Expect(t *testing.T, actual interface{}) *ProxyTestee {
	return &ProxyTestee{t: t, actual: actual}
}
func (p *ProxyTestee) failed(fail ...int) {
	f := FailBase
	if 0 < len(fail) {
		f = fail[0]
	}
	p.failWith(f)
}
func (p *ProxyTestee) failWith(fail int) {
	fmt.Printf(
		Scolds[fail],
		p.expected,
		p.actual,
	)
	p.t.Fail()
	os.Exit(1)
}
func (p *ProxyTestee) ToBe(expected interface{}) {
	if p.actual == expected {
		return
	}
	p.expected = expected
	p.failed()
}
