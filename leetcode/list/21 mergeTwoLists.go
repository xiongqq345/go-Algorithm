package list

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	start := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			dummy.Next = l1
			l1 = l1.Next
		} else {
			dummy.Next = l2
			l2 = l2.Next
		}
		dummy = dummy.Next
	}
	if l1 == nil {
		dummy.Next = l2
	} else {
		dummy.Next = l1
	}
	return start.Next
}
