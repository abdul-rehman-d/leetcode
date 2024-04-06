package is_palindrome_test

import (
	"testing"

	"github.com/abdul-rehman-d/leetcode/internal/is_palindrome"
)

func TestMain(t *testing.T) {
	inputs := [3]string {
		"A man, a plan, a canal: Panama",
		"race a car",
		" ",
	}
	outputs := [3]bool{true, false, true}

	for idx, input := range inputs {
		result := is_palindrome.IsPalindrome(input)
		if result != outputs[idx] {
			t.Logf("\nInput: %v\nExpected: %v\nActual: %v", input, outputs[idx], result)
			t.Fail()
		}
	}
}

