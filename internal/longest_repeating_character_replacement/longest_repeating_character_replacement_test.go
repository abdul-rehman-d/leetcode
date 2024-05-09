package longest_repeating_character_replacement_test

import (
	"testing"

	"github.com/abdul-rehman-d/leetcode/internal/longest_repeating_character_replacement"
)

func TestMain(t *testing.T) {
	inputs := [3]struct {
		s string
		k int
	}{
		{s: "ABAB", k: 2},
		{s: "AABABBA", k: 1},
		{s: "AAAA", k: 0},
	}
	outputs := [3]int{4, 4, 4}

	for idx, input := range inputs {
		result := longest_repeating_character_replacement.CharacterReplacement(input.s, input.k)
		if result != outputs[idx] {
			t.Logf("\nInput: %v\nExpected: %v\nActual: %v", input, outputs[idx], result)
			t.Fail()
		}
	}
}
