package linked_list

// 给出一个升序排序的链表，删除链表中的所有重复出现的元素，只保留原链表中只出现一次的元素。

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val != cur.Next.Next.Val {
			cur = cur.Next
		} else {
			v := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == v {
				cur.Next = cur.Next.Next
			}
		}
	}
	return dummy.Next
}
