package reverse_linked_list

import (
	"testing"

	"github.com/abdul-rehman-d/leetcode/internal/common"
)

type TestCase struct {
	Input  []int
	Output []int
}

func TestMain(t *testing.T) {
	testCases := []TestCase{
		{Input: []int{}, Output: []int{}},
		{Input: []int{1, 2}, Output: []int{2, 1}},
		{Input: []int{1, 2, 3, 4, 5}, Output: []int{5, 4, 3, 2, 1}},
	}
	for _, tc := range testCases {
		i, o := common.NewListNode(tc.Input), common.NewListNode(tc.Output)
		result := ReverseLinkedList(i)
		if !result.Compare(o) {
			t.Logf(
				"\nInput: %+v\nExpected: %+v\nActual: %+v",
				i.ToString(),
				o.ToString(),
				result.ToString(),
			)
			t.Fail()
		}
	}
}
