package mint

import "reflect"

func (testee *Testee) ToBe(expected interface{}) *Testee {
	if testee.actual == expected {
		return testee
	}
	testee.expected = expected
	return testee.failed()
}

// FIXME: Is `string` the base way?
func (testee *Testee) TypeOf(typeName string) *Testee {
	if reflect.TypeOf(testee.actual).String() == typeName {
		return testee
	}
	testee.expected = typeName
	return testee.failed()
}
