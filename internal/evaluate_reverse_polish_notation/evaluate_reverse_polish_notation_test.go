package evaluate_reverse_polish_notation

import (
	"testing"
)

type TestCase struct {
	Input  []string
	Output int
}

func TestMain(t *testing.T) {
	testCases := []TestCase{
		{[]string{"2", "1", "+", "3", "*"}, 9},
		{[]string{"4", "13", "5", "/", "+"}, 6},
		{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}
	for _, tc := range testCases {
		result := EvaluateReversePolishNotation(tc.Input)
		if result != tc.Output {
			t.Logf("\nInput: %+v\nExpected: %+v\nActual: %+v", tc.Input, tc.Output, result)
			t.Fail()
		}
	}
}
