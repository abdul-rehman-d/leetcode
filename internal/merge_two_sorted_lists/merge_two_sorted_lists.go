package merge_two_sorted_lists

import (
	"github.com/abdul-rehman-d/leetcode/internal/common"
)

type ListNode = common.ListNode

func MergeTwoSortedLists(list1 *ListNode, list2 *ListNode) *ListNode {
	curr := &ListNode{}
	head := curr
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}
	if list1 != nil {
		curr.Next = list1
	}
	if list2 != nil {
		curr.Next = list2
	}

	return nil
}
