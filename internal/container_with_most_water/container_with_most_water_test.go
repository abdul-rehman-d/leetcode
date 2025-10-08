package container_with_most_water_test

import (
	"testing"

	"github.com/abdul-rehman-d/leetcode/internal/container_with_most_water"
)

func TestMain(t *testing.T) {
	inputs := [2][]int{
		{1, 8, 6, 2, 5, 4, 8, 3, 7},
		{1, 1},
	}
	outputs := [2]int{49, 1}

	for idx, input := range inputs {
		result := container_with_most_water.MaxArea(input)
		if result != outputs[idx] {
			t.Logf("\nInput: %v\nExpected: %v\nActual: %v", input, outputs[idx], result)
			t.Fail()
		}
	}
}
