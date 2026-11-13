package reverse_linked_list

import "github.com/abdul-rehman-d/leetcode/internal/common"

type ListNode = common.ListNode

func ReverseLinkedList(head *ListNode) *ListNode {
	var recurse func(*ListNode) *ListNode
	recurse = func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}
		newHead := recurse(head.Next)
		head.Next.Next = head
		head.Next = nil
		return newHead
	}

	return recurse(head)
}
