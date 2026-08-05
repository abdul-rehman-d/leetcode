package find_minimum_in_rotated_sorted_array

import (
	"testing"
)

type TestCase struct {
	Input  []int
	Output int
}

func TestMain(t *testing.T) {
	testCases := []TestCase{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{11, 13, 15, 17}, 11},
		{[]int{2, 1}, 1},
		{[]int{3, 1, 2}, 1},
		{[]int{1, 2, 3, 4, 5, 6}, 1},
		{[]int{3, 4, 5, 6, 1, 2}, 1},
		{[]int{6, 1, 2, 3, 4, 5}, 1},
	}
	for _, tc := range testCases {
		result := FindMinimumInRotatedSortedArray(tc.Input)
		if result != tc.Output {
			t.Logf("\nInput: %+v\nExpected: %+v\nActual: %+v", tc.Input, tc.Output, result)
			t.Fail()
		}
	}
}
