package linked_list

//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	p := dummy
	for p.Next != nil && p.Next.Next != nil {
		node1 := p.Next
		node2 := p.Next.Next
		p.Next, node1.Next, node2.Next = node2, node2.Next, node1
		p = node1
	}
	return dummy.Next
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(head.Next.Next)
	newHead.Next = head
	return newHead
}
