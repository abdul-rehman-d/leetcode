package permutation_in_string_test

import (
	"testing"

	"github.com/abdul-rehman-d/leetcode/internal/permutation_in_string"
)

func TestMain(t *testing.T) {
	inputs := [4][2]string{
		{"ab", "eidbaooo"},
		{"ab", "eidboaoo"},
		{"abc", "dcba"},
		{"abc", "bbbca"},
	}
	outputs := [4]bool{true, false, true, true}

	for idx, input := range inputs {
		result := permutation_in_string.CheckInclusion(input[0], input[1])
		if result != outputs[idx] {
			t.Logf("\nInput: %v\nExpected: %v\nActual: %v", input, outputs[idx], result)
			t.Fail()
		}
	}
}
