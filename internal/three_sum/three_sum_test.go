package three_sum_test

import (
	"testing"

	"github.com/abdul-rehman-d/leetcode/internal/three_sum"
)

func same2(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func same1(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if !same2(a[i], b[i]) {
			return false
		}
	}
	return true
}

func TestMain(t *testing.T) {
	inputs := [3][]int{
		{-1, 0, 1, 2, -1, -4},
		{0, 1, 1},
		{0, 0, 0},
	}
	outputs := [3][][]int{
		{{-1, -1, 2}, {-1, 0, 1}},
		{},
		{{0, 0, 0}},
	}

	for idx, input := range inputs {
		actual := three_sum.ThreeSum(input)

		if !same1(actual, outputs[idx]) {
			t.Logf("\nInput: %v\nExpected: %v\nActual: %v", input, outputs[idx], actual)
			t.Fail()
		}
	}
}
