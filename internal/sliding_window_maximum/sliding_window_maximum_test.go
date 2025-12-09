package sliding_window_maximum

import (
	"slices"
	"testing"
)

func TestMain(t *testing.T) {
	inputs := []struct {
		arr []int
		k   int
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3},
		{[]int{1, 3, -1, -3, 0, 3, 6, 7}, 3},
		{[]int{1}, 1},
		{[]int{1, -1}, 1},
	}
	outputs := [][]int{
		{3, 3, 5, 5, 6, 7},
		{3, 3, 0, 3, 6, 7},
		{1},
		{1, -1},
	}

	for idx, input := range inputs {
		result := MaxSlidingWindow(input.arr, input.k)
		if slices.Compare(result, outputs[idx]) != 0 {
			t.Logf("\nInput: %v\nExpected: %v\nActual: %v", input, outputs[idx], result)
			t.Fail()
		}
	}
}
