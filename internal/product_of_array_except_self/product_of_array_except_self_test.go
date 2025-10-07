package product_of_array_except_self

import (
	"slices"
	"testing"
)

func TestMain(t *testing.T) {
	inputs := [2][]int{
		{1, 2, 3, 4},
		{-1, 1, 0, -3, 3},
	}
	outputs := [2][]int{
		{24, 12, 8, 6},
		{0, 0, 9, 0, 0},
	}

	for idx, input := range inputs {
		result := ProductExceptSelf(input)
		if slices.Compare(result, outputs[idx]) != 0 {
			t.Logf("\nInput: %v\nExpected: %v\nActual: %v", input, outputs[idx], result)
			t.Fail()
		}
	}
}
