package two_sum_ii_test

import (
	"testing"

	"github.com/abdul-rehman-d/leetcode/internal/two_sum_ii"
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
	inputs := [4]struct{
		Arr []int;
		Target int;
	}{
		{Arr: []int{2,7,11,15}, Target: 9},
		{Arr: []int{2,3,4}, Target: 6},
		{Arr: []int{-1,0}, Target: -1},
		{Arr: []int{5,25,75}, Target: 100},
	}
	outputs := [4][]int{
		{1,2},
		{1,3},
		{1,2},
		{2,3},
	}

	for idx, input := range inputs {
		result := two_sum_ii.TwoSum(input.Arr, input.Target)
		if !areArrsSame(result, outputs[idx]) {
			t.Logf("\nInput: %v\nExpected: %v\nActual: %v", input, outputs[idx], result)
			t.Fail()
		}
	}
}


