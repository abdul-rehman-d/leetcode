package is_anagram_test

import (
	"testing"

	"github.com/abdul-rehman-d/leetcode/internal/is_anagram"
)

func TestMain(t *testing.T) {
	inputs := [2][2]string{
		{"anagram", "nagaram"},
		{"rat", "car"},
	}
	outputs := [2]bool{true, false}

	for idx, input := range inputs {
		result := is_anagram.IsAnagram(input[0], input[1])
		if result != outputs[idx] {
			t.Logf("\nInput: %v\nExpected: %v\nActual: %v", input, outputs[idx], result)
			t.Fail()
		}
	}
}
