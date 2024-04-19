package longest_substring_without_repeating_characters_test

import (
	"testing"

	"github.com/abdul-rehman-d/leetcode/internal/longest_substring_without_repeating_characters"
)

func TestMain(t *testing.T) {
	inputs := [4]string{
		"abcabcbb",
		"bbbbb",
		"pwwkew",
        "dvdf",
	}
	outputs := [4]int{3, 1, 3,3}

	for idx, input := range inputs {
		result := longest_substring_without_repeating_characters.LengthOfLongestSubstring(input)
		if result != outputs[idx] {
			t.Logf("\nInput: %v\nExpected: %v\nActual: %v", input, outputs[idx], result)
			t.Fail()
		}
	}
}
