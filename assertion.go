package mint

import "reflect"

func (testee *Testee) ToBe(expected interface{}) *Testee {
	if judge(testee.actual, expected, testee.not) {
		return testee
	}
	testee.expected = expected
	return testee.failed()
}

// FIXME: Is `string` the base way?
func (testee *Testee) TypeOf(typeName string) *Testee {
	if judge(reflect.TypeOf(testee.actual).String(), typeName, testee.not) {
		return testee
	}
	testee.expected = typeName
	return testee.failed()
}
