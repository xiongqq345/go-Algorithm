package linked_list

// K 个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	p := head
	for i := 0; i < k; i++ {
		if p == nil {
			return head
		}
		p = p.Next
	}
	res := reverseToNode(head, p)
	head.Next = reverseKGroup(p, k)
	return res
}
