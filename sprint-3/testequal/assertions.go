//go:build !solution

package testequal

import (
	"fmt"
	"reflect"
)

func compare(e, a interface{}) bool {
	switch e.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return (reflect.TypeOf(e) == reflect.TypeOf(a)) && (e == a)
	case string, map[string]string, []int, []byte:
		return reflect.DeepEqual(e, a)
	default:
		return false
	}

}
func assert(t T, expected, actual interface{}, negate bool, msgAndArgs ...interface{}) bool {
	t.Helper()
	if isEq := compare(expected, actual); (isEq && !negate) || (!isEq && negate) {
		fmt.Println(isEq)
		return true
	}

	if len(msgAndArgs) == 0 {
		t.Errorf("")
		return false
	}

	if format, ok := msgAndArgs[0].(string); ok {
		t.Errorf(format, msgAndArgs[1:]...)
	} else {
		t.Errorf("", msgAndArgs)
	}

	return false
}

// AssertEqual checks that expected and actual are equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are equal.
func AssertEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	t.Helper()
	return assert(t, expected, actual, false, msgAndArgs...)
}

// AssertNotEqual checks that expected and actual are not equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are not equal.
func AssertNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	t.Helper()
	return assert(t, expected, actual, true, msgAndArgs...)
}

// RequireEqual does the same as AssertEqual but fails caller test immediately.
func RequireEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	if !AssertEqual(t, expected, actual, msgAndArgs...) {
		t.FailNow()
	}
}

// RequireNotEqual does the same as AssertNotEqual but fails caller test immediately.
func RequireNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	if !AssertNotEqual(t, expected, actual, msgAndArgs...) {
		t.FailNow()
	}
}
