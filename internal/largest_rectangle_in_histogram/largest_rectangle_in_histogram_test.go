package largest_rectangle_in_histogram

import (
	"testing"
)

type TestCase struct {
	Input  []int
	Output int
}

func TestMain(t *testing.T) {
	testCases := []TestCase{
		{Input: []int{2, 1, 5, 6, 2, 3}, Output: 10},
		{Input: []int{2, 4}, Output: 4},
		{Input: []int{7, 1, 7, 2, 2, 4}, Output: 8},
		{Input: []int{1, 3, 7}, Output: 7},
		{Input: []int{3, 6, 5, 7, 4, 8, 1}, Output: 20},
	}
	for _, tc := range testCases {
		result := LargestRectangleInHistogram(tc.Input)
		if result != tc.Output {
			t.Logf("\nInput: %+v\nExpected: %+v\nActual: %+v", tc.Input, tc.Output, result)
			t.Fail()
		}
	}
}
