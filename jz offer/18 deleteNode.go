package jz_offer

// 给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
//
//返回删除后的链表的头节点。
func deleteNode(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	node := dummy
	for node.Next != nil {
		if node.Next.Val == val {
			node.Next = node.Next.Next
			break
		}
		node = node.Next
	}
	return dummy.Next
}
