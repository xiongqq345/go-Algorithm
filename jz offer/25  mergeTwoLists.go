package jz_offer

// 输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	p := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 == nil {
		p.Next = l2
	} else {
		p.Next = l1
	}
	return dummy.Next
}
