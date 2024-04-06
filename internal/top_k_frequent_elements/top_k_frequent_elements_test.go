package top_k_frequent_elements_test

import (
	"testing"

	"github.com/abdul-rehman-d/leetcode/internal/top_k_frequent_elements"
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
	inputs := [2]struct{
		Arr []int;
		K int;
	}{
		{Arr: []int{1,1,1,2,2,3}, K: 2},
		{Arr: []int{1}, K: 1},
	}
	outputs := [2][]int{
		{1,2},
		{1},
	}

	for idx, input := range inputs {
		result := top_k_frequent_elements.TopKFrequent(input.Arr, input.K)
		if !areArrsSame(result, outputs[idx]) {
			t.Logf("\nInput: %v\nExpected: %v\nActual: %v", input, outputs[idx], result)
			t.Fail()
		}
	}
}

