package linked_list

//给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
//
//返回删除后的链表的头节点。
func deleteNode(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	pre, cur := dummy, head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			break
		}
		cur, pre = cur.Next, cur
	}
	return dummy.Next
}
