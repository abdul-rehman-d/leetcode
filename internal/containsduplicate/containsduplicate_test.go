package containsduplicate_test

import (
	"testing"

	"github.com/abdul-rehman-d/leetcode/internal/containsduplicate"
)

func TestMain(t *testing.T) {
	inputs := [3][]int{
		{1,2,3,1},
		{1,2,3,4},
		{1,1,1,3,3,4,3,2,4,2},
	}
	outputs := [3]bool{true, false, true}

	for idx, input := range inputs {
		result := containsduplicate.ContainsDuplicate(input)
		if result != outputs[idx] {
			t.Logf("\nInput: %v\nExpected: %v\nActual: %v", input, outputs[idx], result)
			t.Fail()
		}
	}
}

