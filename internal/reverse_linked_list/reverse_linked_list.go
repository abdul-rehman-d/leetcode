package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseLinkedList(head *ListNode) *ListNode {
	/*
		i=0
		temp = nil
		y    = 2->3->4->5
		x    = 1->nil
		temp = 1->nil
		x    = 2->3->4->5
		y    = 1->nil
		i=1
		temp = 1->nil
		y    = 3->4->5
		x    = 2->1->nil
		temp = 2->1->nil
		x    = 3->4->5
		y    = 2->1->nil
		i=2
		temp = 2->1->nil
		y    = 4->5
		x    = 3->2->1->nil
		temp = 3->2->1->nil
		x    = 4->5
		y    = 3->2->1->nil
		i=3
		temp = 3->2->1->nil
		y    = 5->nil
		x    = 4->3->2->1->nil
		temp = 4->3->2->1->nil
		x    = 5->nil
		y    = 4->3->2->1->nil
		i=4
		temp = 4->3->2->1->nil
		y    = nil
		x    = 5->4->3->2->1->nil
		temp = 5->4->3->2->1->nil
		x    = nil
		y    = 5->4->3->2->1->nil

	*/

	if head == nil {
		return nil
	}

	var prev *ListNode = nil

	for head != nil {
		prev, head.Next = head.Next, prev
		head, prev = prev, head
	}

	return prev
}
