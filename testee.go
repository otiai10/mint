package mint

import "testing"
import "reflect"
import "fmt"
import "runtime"
import "path/filepath"

// Testee is holder of interfaces which user want to assert
// and also has its result.
type Testee struct {
	t        *testing.T
	actual   interface{}
	expected interface{}
	dry      bool
	not      bool
	deeply   bool
	result   Result
}

// ToBe can assert the testee to equal the parameter of this func.
// OS will exit with code 1, when the assertion fail.
// If you don't want to exit, see "Dry()".
func (testee *Testee) ToBe(expected interface{}) Result {
	if judge(testee.actual, expected, testee.not, testee.deeply) {
		return testee.result
	}
	testee.expected = expected
	return testee.failed(failToBe)
}

// In can assert the testee is in given array.
func (testee *Testee) In(expecteds ...interface{}) Result {
	for _, expected := range expecteds {
		if judge(testee.actual, expected, testee.not, testee.deeply) {
			return testee.result
		}
	}
	testee.expected = expecteds
	return testee.failed(failIn)
}

// TypeOf can assert the type of testee to equal the parameter of this func.
// OS will exit with code 1, when the assertion fail.
// If you don't want to exit, see "Dry()".
func (testee *Testee) TypeOf(typeName string) Result {
	if judge(reflect.TypeOf(testee.actual).String(), typeName, testee.not, testee.deeply) {
		return testee.result
	}
	testee.expected = typeName
	return testee.failed(failType)
}

// Not makes following assertion conversed.
func (testee *Testee) Not() *Testee {
	testee.not = true
	return testee
}

// Dry makes the testee NOT to call "Fail()".
// Use this if you want to fail test in a purpose.
func (testee *Testee) Dry() *Testee {
	testee.dry = true
	return testee
}

// Deeply makes following assertions use `reflect.DeepEqual`.
// You had better use this to compare reference type objects.
func (testee *Testee) Deeply() *Testee {
	testee.deeply = true
	return testee
}

func (testee *Testee) failed(failure int) Result {
	message := testee.toText(failure)
	testee.result.ok = false
	testee.result.message = message
	if !testee.dry {
		fmt.Println(colorize["red"](message))
		testee.t.Fail()
	}
	return testee.result
}
func (testee *Testee) toText(fail int) string {
	not := ""
	if testee.not {
		not = "NOT "
	}
	_, file, line, _ := runtime.Caller(3)
	return fmt.Sprintf(
		scolds[fail],
		filepath.Base(file), line,
		not,
		testee.expected,
		testee.actual,
	)
}
