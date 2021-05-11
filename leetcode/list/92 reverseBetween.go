package list

// 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var l, r, pre, succ, p *ListNode
	dummy := &ListNode{Next: head}
	pre = dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	l, r = pre.Next, pre.Next

	for i := 0; i < right-left; i++ {
		r = r.Next
	}
	succ = r.Next

	c := l
	for c != succ {
		c, c.Next, p = c.Next, p, c
	}
	if pre != nil {
		pre.Next = r
	}
	l.Next = succ
	return dummy.Next
}
