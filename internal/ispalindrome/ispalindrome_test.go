package ispalindrome_test

import (
	"testing"

	"github.com/abdul-rehman-d/leetcode/internal/ispalindrome"
)

func TestMain(t *testing.T) {
	inputs := [3]string {
		"A man, a plan, a canal: Panama",
		"race a car",
		" ",
	}
	outputs := [3]bool{true, false, true}

	for idx, input := range inputs {
		result := ispalindrome.IsPalindrome(input)
		if result != outputs[idx] {
			t.Logf("\nInput: %v\nExpected: %v\nActual: %v", input, outputs[idx], result)
			t.Fail()
		}
	}
}

