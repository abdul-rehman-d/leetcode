package contains_duplicate_test

import (
	"testing"

	"github.com/abdul-rehman-d/leetcode/internal/contains_duplicate"
)

func TestMain(t *testing.T) {
	inputs := [3][]int{
		{1,2,3,1},
		{1,2,3,4},
		{1,1,1,3,3,4,3,2,4,2},
	}
	outputs := [3]bool{true, false, true}

	for idx, input := range inputs {
		result := contains_duplicate.ContainsDuplicate(input)
		if result != outputs[idx] {
			t.Logf("\nInput: %v\nExpected: %v\nActual: %v", input, outputs[idx], result)
			t.Fail()
		}
	}
}

