package mint

import "reflect"

func (p *ProxyTestee) ToBe(expected interface{}) *ProxyTestee {
	if p.actual == expected {
		return p
	}
	p.expected = expected
	return p.failed()
}

// FIXME: Is `string` the base way?
func (p *ProxyTestee) TypeOf(typeName string) *ProxyTestee {
	if reflect.TypeOf(p.actual).String() == typeName {
		return p
	}
	p.expected = typeName
	return p.failed()
}
