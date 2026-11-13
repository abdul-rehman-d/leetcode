package median_of_two_sorted_arrays

import (
	"testing"
)

type TestCase struct {
	Input1 []int
	Input2 []int
	Output float64
}

func TestMain(t *testing.T) {
	testCases := []TestCase{
		{[]int{1}, []int{}, 1.0},
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5, 6, 7, 8}, 4.0},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4, 5, 6, 7, 8}, 3.5},
		{[]int{1, 2, 3, 4, 5}, []int{}, 3.0},
		{[]int{5, 6}, []int{1, 2, 3, 4, 7, 8}, 4.5},
	}
	for _, tc := range testCases {
		result := MedianOfTwoSortedArrays(tc.Input1, tc.Input2)
		if result != tc.Output {
			t.Logf("\nInput1: %+v\nInput2: %+v\nExpected: %+v\nActual: %+v", tc.Input1, tc.Input2, tc.Output, result)
			t.Fail()
		}
	}
}
