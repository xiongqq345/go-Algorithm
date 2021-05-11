package list

// 输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。
func getKthFromEnd(head *ListNode, k int) *ListNode {
	p1, p2 := head, head
	for i := 0; i < k; i++ {
		p2 = p2.Next
	}
	for p2 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}
