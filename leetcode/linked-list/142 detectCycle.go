package linked_list

// 给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]int)
	for head != nil {
		if _, ok := m[head]; ok {
			return head
		}
		m[head] = 0
		head = head.Next
	}
	return nil
}
