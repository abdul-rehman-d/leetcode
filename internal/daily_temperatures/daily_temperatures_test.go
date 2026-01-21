package daily_temperatures

import (
	"slices"
	"testing"
)

type TestCase struct {
	Input  []int
	Output []int
}

func TestMain(t *testing.T) {
	testCases := []TestCase{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{[]int{30, 60, 90}, []int{1, 1, 0}},
	}
	for _, tc := range testCases {
		result := DailyTemperatures(tc.Input)
		if slices.Compare(result, tc.Output) != 0 {
			t.Logf("\nInput: %+v\nExpected: %+v\nActual: %+v", tc.Input, tc.Output, result)
			t.Fail()
		}
	}
}
