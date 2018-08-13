package mint

import (
	"os"

	"github.com/bouk/monkey"
)

// Exit ...
func (testee *Testee) Exit(expectedCode int) Result {

	fun, ok := testee.actual.(func())
	if !ok {
		panic("mint error: Exit only can be called for func type value")
	}

	var actualCode int
	patch := monkey.Patch(os.Exit, func(code int) {
		actualCode = code
	})
	fun()
	patch.Unpatch()

	testee.actual = actualCode
	if judge(actualCode, expectedCode, testee.not, testee.deeply) {
		return testee.result
	}
	testee.expected = expectedCode
	return testee.failed(failExitCode)
}
