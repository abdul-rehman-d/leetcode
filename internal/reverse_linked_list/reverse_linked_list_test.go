package reverse_linked_list

import (
	"testing"
)

type TestCase struct {
	Input  *ListNode
	Output *ListNode
}

func compareLinkedLists(a, b *ListNode) bool {
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return a == nil && b == nil
}

func TestMain(t *testing.T) {
	testCases := []TestCase{
		{Input: nil, Output: nil},
		{
			Input:  &ListNode{1, &ListNode{2, nil}},
			Output: &ListNode{2, &ListNode{1, nil}},
		},
		{
			Input:  &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			Output: &ListNode{5, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}},
		},
	}
	for _, tc := range testCases {
		result := ReverseLinkedList(tc.Input)
		if !compareLinkedLists(result, tc.Output) {
			t.Logf("\nInput: %+v\nExpected: %+v\nActual: %+v", tc.Input, tc.Output, result)
			t.Fail()
		}
	}
}
