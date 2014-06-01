package mint

import "reflect"
import "fmt"
import "os"

// "*Testee.ToBe" can assert the testee to equal the parameter of this func.
// OS will exit with code 1, when the assertion fail.
// If you don't want to exit, see "Dry()".
func (testee *Testee) ToBe(expected interface{}) *Testee {
	if judge(testee.actual, expected, testee.not) {
		return testee
	}
	testee.expected = expected
	return testee.failed(FailBase)
}

// "*Testee.TypeOf" can assert the type of testee to equal the parameter of this func.
// OS will exit with code 1, when the assertion fail.
// If you don't want to exit, see "Dry()".
func (testee *Testee) TypeOf(typeName string) *Testee {
	if judge(reflect.TypeOf(testee.actual).String(), typeName, testee.not) {
		return testee
	}
	testee.expected = typeName
	return testee.failed(FailType)
}

// "*Testee.Not" makes following assertion conversed.
func (testee *Testee) Not() *Testee {
	testee.not = true
	return testee
}

// "*Testee.Dry" makes the testee NOT to call "os.Exit(1)".
// Use this if you want to fail test in a purpose.
func (testee *Testee) Dry() *Testee {
	testee.dry = true
	return testee
}

func (testee *Testee) failed(failure int) *Testee {
	message := testee.toText(failure)
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
