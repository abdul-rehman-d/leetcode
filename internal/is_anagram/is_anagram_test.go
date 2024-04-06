package isanagram_test

import (
	"testing"

	"github.com/abdul-rehman-d/leetcode/internal/isanagram"
)

func TestMain(t *testing.T) {
	inputs := [2][2]string{
		{"anagram", "nagaram"},
		{"rat", "car"},
	}
	outputs := [2]bool{true, false}

	for idx, input := range inputs {
		result := isanagram.IsAnagram(input[0], input[1])
		if result != outputs[idx] {
			t.Logf("\nInput: %v\nExpected: %v\nActual: %v", input, outputs[idx], result)
			t.Fail()
		}
	}
}

