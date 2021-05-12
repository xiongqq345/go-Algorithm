package list

// 存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素 只出现一次 。
//
//返回同样按升序排列的结果链表。

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	pre, cur := head, head.Next
	for cur != nil {
		if pre.Val == cur.Val {
			pre.Next = cur.Next
			cur = cur.Next
		} else {
			pre, cur = cur, cur.Next
		}
	}
	return head
}
