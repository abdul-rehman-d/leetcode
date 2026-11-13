package reverse_linked_list

import (
	"strconv"
	"strings"
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

func printLinkedList(a *ListNode) string {
	var result strings.Builder
	seen := make(map[*ListNode]bool)
	for a != nil {
		if result.Len() > 0 {
			result.WriteString("->")
		}
		if seen[a] {
			result.WriteString("(cycle)")
			break
		}
		seen[a] = true
		result.WriteString(strconv.Itoa(a.Val))
		a = a.Next
	}
	return result.String()
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
			t.Logf(
				"\nInput: %+v\nExpected: %+v\nActual: %+v",
				printLinkedList(tc.Input),
				printLinkedList(tc.Output),
				printLinkedList(result),
			)
			t.Fail()
		}
	}
}
