package linked_list

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	p1, p2 := dummy, dummy
	for n > 0 {
		p2 = p2.Next
		n--
	}
	for p2.Next != nil {
		p1, p2 = p1.Next, p2.Next
	}
	p1.Next = p1.Next.Next
	return dummy.Next
}
