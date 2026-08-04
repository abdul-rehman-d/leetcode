package core

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		Arr    []int
		Target int
		Output int
	}{
		{[]int{1, 2, 3, 4, 5}, 1, 0},
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{1, 2, 3, 4, 5}, 5, 4},
		{[]int{2, 4, 6, 8, 10, 12}, 8, 3},
		{[]int{2, 4, 6, 8, 10, 12}, 7, -1},
		{[]int{-1, 0, 3, 5, 9, 12}, 13, -1},
	}

	for _, test := range tests {
		result := BinarySearch(test.Arr, test.Target)
		if result != test.Output {
			t.Logf("\nInput: %v\t%v\nExpected: %v\nActual: %v", test.Arr, test.Target, test.Output, result)
			t.Fail()
		}
	}
}
