package mint

import "reflect"
import "fmt"
import "os"

// ToBe can assert the testee to equal the parameter of this func.
// OS will exit with code 1, when the assertion fail.
// If you don't want to exit, see "Dry()".
func (testee *Testee) ToBe(expected interface{}) Result {
	if judge(testee.actual, expected, testee.not, testee.deeply) {
		return testee.Result
	}
	testee.expected = expected
	return testee.failed(failToBe)
}

// TypeOf can assert the type of testee to equal the parameter of this func.
// OS will exit with code 1, when the assertion fail.
// If you don't want to exit, see "Dry()".
func (testee *Testee) TypeOf(typeName string) Result {
	if judge(reflect.TypeOf(testee.actual).String(), typeName, testee.not, testee.deeply) {
		return testee.Result
	}
	testee.expected = typeName
	return testee.failed(failType)
}

// Not makes following assertion conversed.
func (testee *Testee) Not() *Testee {
	testee.not = true
	return testee
}

// Dry makes the testee NOT to call "os.Exit(1)".
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
	if !testee.dry {
		fmt.Println(colorize["red"](message))
		testee.t.Fail()
		os.Exit(1)
	}
	testee.Result.ok = false
	testee.Result.message = message
	return testee.Result
}
func (testee *Testee) toText(fail int) string {
	not := ""
	if testee.not {
		not = "NOT "
	}
	return fmt.Sprintf(
		scolds[fail],
		not,
		testee.expected,
		testee.actual,
	)
}
