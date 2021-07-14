package linked_list

// 存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素 只出现一次 。
//
//返回同样按升序排列的结果链表。

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	p := head
	for p != nil {
		next := p.Next
		for next != nil && next.Val == p.Val {
			next = next.Next
		}
		p, p.Next = next, next
	}
	return head
}
