package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	new_link_head := ListNode{
		Val:  0,
		Next: nil,
	}
	new_head_pointer := &new_link_head

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			new_head_pointer.Next = l1
			l1 = l1.Next
		} else {
			new_head_pointer.Next = l2
			l2 = l2.Next
		}
		new_head_pointer = new_head_pointer.Next
	}

	if l1 == nil {
		new_head_pointer.Next = l2
	}

	if l2 == nil {
		new_head_pointer.Next = l1
	}
	return new_link_head.Next
}
