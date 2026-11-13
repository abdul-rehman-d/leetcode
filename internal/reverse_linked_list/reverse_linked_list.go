package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseLinkedList(head *ListNode) *ListNode {
	var recurse func(*ListNode) *ListNode
	recurse = func(head *ListNode) *ListNode {
		if head == nil {
			return nil
		}
		x := head
		if head.Next != nil {
			x = recurse(head.Next)
			head.Next.Next = head
		}
		head.Next = nil
		return x
	}

	return recurse(head)
}
