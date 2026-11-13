package merge_two_sorted_lists

import (
	"testing"

	"github.com/abdul-rehman-d/leetcode/internal/common"
)

type TestCase struct {
	Input1 []int
	Input2 []int
	Output []int
}

func TestMain(t *testing.T) {
	testCases := []TestCase{
		{Input1: []int{}, Input2: []int{}, Output: []int{}},
		{Input1: []int{}, Input2: []int{0}, Output: []int{0}},
		{Input1: []int{1, 2, 4}, Input2: []int{1, 3, 4}, Output: []int{1, 1, 2, 3, 4, 4}},
	}
	for _, tc := range testCases {
		i1 := common.NewListNode(tc.Input1)
		i2 := common.NewListNode(tc.Input2)
		o := common.NewListNode(tc.Output)

		result := MergeTwoSortedLists(i1, i2)

		if !result.Compare(o) {
			t.Logf(
				"\nInput1: %+v\nInput2: %+v\nExpected: %+v\nActual: %+v",
				i1.ToString(),
				i2.ToString(),
				o.ToString(),
				result.ToString(),
			)
			t.Fail()
		}
	}
}
