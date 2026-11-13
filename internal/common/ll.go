package common

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (a *ListNode) ToString() string {
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

func (a *ListNode) Compare(b *ListNode) bool {
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return a == nil && b == nil
}

func NewListNode(rawList []int) *ListNode {
	if len(rawList) == 0 {
		return nil
	}

	head := &ListNode{}
	x := head

	for i, v := range rawList {
		x.Val = v
		if i != len(rawList)-1 {
			x.Next = &ListNode{}
			x = x.Next
		}
	}

	return head
}
