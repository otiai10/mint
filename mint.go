package mint

import "testing"
import "fmt"
import "os"

type ProxyTestee struct {
    t *testing.T
    actual interface{}
}

func Expect(t *testing.T, actual interface{}) *ProxyTestee {
    return &ProxyTestee{t, actual}
}
func (p *ProxyTestee) ToBe(expected interface{}) {
    if p.actual == expected {
        return
    }
	fmt.Printf("Expected to be `%+v`, but actual `%+v`\n", expected, p.actual)
    p.t.Fail()
    os.Exit(1)
}
