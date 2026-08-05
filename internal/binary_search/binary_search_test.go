package binary_search

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
		{[]int{1, 2, 3, 4, 5}, 1, 0},
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{1, 2, 3, 4, 5}, 5, 4},
		{[]int{2, 4, 6, 8, 10, 12}, 8, 3},
		{[]int{2, 4, 6, 8, 10, 12}, 7, -1},
		{[]int{-1, 0, 3, 5, 9, 12}, 13, -1},
	}
	for _, tc := range testCases {
		result := BinarySearch(tc.Input, tc.Target)
		if result != tc.Output {
			t.Logf("\nInput: %+v\nExpected: %+v\nActual: %+v", tc.Input, tc.Output, result)
			t.Fail()
		}
	}
}
