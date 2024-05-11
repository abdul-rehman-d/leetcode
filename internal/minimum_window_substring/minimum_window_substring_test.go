package minimum_window_substring_test

import (
	"testing"

	"github.com/abdul-rehman-d/leetcode/internal/minimum_window_substring"
)


func TestMain(t *testing.T) {
	inputs := [3][2]string{
		{"ADOBECODEBANC", "ABC"},
		{"a", "a"},
		{"aa", "aa"},
	}
	outputs := [3]string{"BANC", "a", "aa"}

	for idx, input := range inputs {
		result := minimum_window_substring.MinWindow(input[0], input[1])
		if result != outputs[idx] {
			t.Logf("\nInput: %v\nExpected: %v\nActual: %v", input, outputs[idx], result)
			t.Fail()
		}
	}
}

