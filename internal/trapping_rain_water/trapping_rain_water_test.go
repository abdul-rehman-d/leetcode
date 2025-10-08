package trapping_rain_water_test

import (
	"testing"

	"github.com/abdul-rehman-d/leetcode/internal/trapping_rain_water"
)

func TestMain(t *testing.T) {
	inputs := [2][]int{
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		{4, 2, 0, 3, 2, 5},
	}
	outputs := [2]int{
		6,
		9,
	}

	for idx, input := range inputs {
		result := trapping_rain_water.Trap(input)
		if result != outputs[idx] {
			t.Logf("\nInput: %v\nExpected: %v\nActual: %v", input, outputs[idx], result)
			t.Fail()
		}
	}
}
