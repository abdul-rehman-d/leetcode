package search_a_2d_matrix

import (
	"testing"
)

type TestCase struct {
	Input  [][]int
	Target int
	Output bool
}

func TestMain(t *testing.T) {
	testCases := []TestCase{
		{
			Input:  [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			Target: 3,
			Output: true,
		},
		{
			Input:  [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			Target: 13,
			Output: false,
		},
		{
			Input:  [][]int{{1, 2, 4, 8}, {10, 11, 12, 13}, {14, 20, 30, 40}},
			Target: 10,
			Output: true,
		},
		{
			Input:  [][]int{{1, 2, 4, 8}, {10, 11, 12, 13}, {14, 20, 30, 40}},
			Target: 15,
			Output: false,
		},
	}
	for _, tc := range testCases {
		result := SearchA2dMatrix(tc.Input, tc.Target)
		if result != tc.Output {
			t.Logf("\nInput: %+v\nExpected: %+v\nActual: %+v", tc.Input, tc.Output, result)
			t.Fail()
		}
	}
}
