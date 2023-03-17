package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	dummy := result
	for l1 != nil || l2 != nil {
		if l1 != nil {
			dummy.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			dummy.Val += l2.Val
			l2 = l2.Next
		}
		if dummy.Val > 9 {
			dummy.Val -= 10
			dummy.Next = &ListNode{Val: 1}
		} else if l1 != nil || l2 != nil {
			dummy.Next = &ListNode{}
		}
		dummy = dummy.Next
	}
	return result
}
