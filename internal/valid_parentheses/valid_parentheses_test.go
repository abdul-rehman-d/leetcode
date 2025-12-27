package valid_parentheses

import (
	"testing"
)

type TestCase struct {
	Input  string
	Output bool
}

func TestMain(t *testing.T) {
	testCases := []TestCase{
		{Input: "()", Output: true},
		{Input: "()[]{}", Output: true},
		{Input: "(]", Output: false},
		{Input: "([])", Output: true},
		{Input: "([)]", Output: false},
	}
	for _, tc := range testCases {
		result := ValidParentheses(tc.Input)
		if result != tc.Output {
			t.Logf("\nInput: %+v\nExpected: %+v\nActual: %+v", tc.Input, tc.Output, result)
			t.Fail()
		}
	}
}

