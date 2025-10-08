package two_sum_test

import (
	"testing"

	"github.com/abdul-rehman-d/leetcode/internal/two_sum"
)

func areArrsSame(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	mapp := make(map[int]int, len(arr1))
	for _, n := range arr1 {
		mapp[n]++
	}
	for _, n := range arr2 {
		if mapp[n] == 0 {
			return false
		}
		mapp[n]--
	}
	for _, val := range mapp {
		if val != 0 {
			return false
		}
	}

	return true
}

func TestMain(t *testing.T) {
	inputs := [3]struct {
		Arr    []int
		Target int
	}{
		{Arr: []int{2, 7, 11, 15}, Target: 9},
		{Arr: []int{3, 2, 4}, Target: 6},
		{Arr: []int{3, 3}, Target: 6},
	}
	outputs := [3][]int{
		{0, 1},
		{1, 2},
		{0, 1},
	}

	for idx, input := range inputs {
		result := two_sum.TwoSum(input.Arr, input.Target)
		if !areArrsSame(result, outputs[idx]) {
			t.Logf("\nInput: %v\nExpected: %v\nActual: %v", input, outputs[idx], result)
			t.Fail()
		}
	}
}
