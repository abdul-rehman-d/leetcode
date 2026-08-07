package search_in_rotated_sorted_array

import (
	"testing"
)

type TestCase struct {
	Input  []int
	Target int
	Output int
}

func TestMain(t *testing.T) {
	testCases := []TestCase{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
		{[]int{3, 4, 5, 6, 1, 2}, 1, 4},
		{[]int{3, 5, 6, 0, 1, 2}, 4, -1},
		{[]int{4, 5, 6, 7, 8, 1, 2, 3}, 8, 4},
	}
	for _, tc := range testCases {
		result := SearchInRotatedSortedArray(tc.Input, tc.Target)
		if result != tc.Output {
			t.Logf("\nInput: %+v, %+v\nExpected: %+v\nActual: %+v", tc.Input, tc.Target, tc.Output, result)
			t.Fail()
		}
	}
}
