package mint

import "reflect"
import "fmt"
import "os"

func (testee *Testee) ToBe(expected interface{}) *Testee {
	if judge(testee.actual, expected, testee.not) {
		return testee
	}
	testee.expected = expected
	return testee.failed(FailBase)
}

// FIXME: Is `string` the base way?
func (testee *Testee) TypeOf(typeName string) *Testee {
	if judge(reflect.TypeOf(testee.actual).String(), typeName, testee.not) {
		return testee
	}
	testee.expected = typeName
	return testee.failed(FailType)
}
func (testee *Testee) Dry() *Testee {
	testee.dry = true
	return testee
}
func (testee *Testee) Not() *Testee {
	testee.not = true
	return testee
}
func (testee *Testee) failed(failure int) *Testee {
	message := testee.toText(failure)
	if !testee.dry {
		testee.t.Errorf(message)
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
