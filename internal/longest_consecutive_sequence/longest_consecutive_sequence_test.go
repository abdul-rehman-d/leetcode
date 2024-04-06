package longestconsecutive_test

import (
	"testing"

	"github.com/abdul-rehman-d/leetcode/internal/longestconsecutive"
)

func TestMain(t *testing.T) {
	inputs := [2][]int {
		{100,4,200,1,3,2},
		{0,3,7,2,5,8,4,6,0,1},
	}
	outputs := [2]int{4, 9}

	for idx, input := range inputs {
		result := longestconsecutive.LongestConsecutive(input)
		if result != outputs[idx] {
			t.Logf("\nInput: %v\nExpected: %v\nActual: %v", input, outputs[idx], result)
			t.Fail()
		}
	}
}


