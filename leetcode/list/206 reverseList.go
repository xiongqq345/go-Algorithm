package list

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

//func reverse(h *ListNode) *ListNode {
//	var p *ListNode
//	for h != nil {
//		h, h.Next, p = h.Next, p, h
//	}
//	return p
//}
