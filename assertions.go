package mint

import "reflect"

func (p *ProxyTestee) ToBe(expected interface{}) {
	if p.actual == expected {
		return
	}
	p.expected = expected
	p.failed()
}

// FIXME: Is `string` the base way?
func (p *ProxyTestee) TypeOf(typeName string) {
	if reflect.TypeOf(p.actual).String() == typeName {
		return
	}
	p.expected = typeName
	p.failed()
}
