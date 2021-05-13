package linked_list

//给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
//将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
//
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

//给定链表 1->2->3->4, 重新排列为 1->4->2->3.

func reorderList(head *ListNode) {
	mid := midNode(head)
	tail := reverse(mid.Next)
	mid.Next = nil
	mergeList(head, tail)
}

func mergeList(n1, n2 *ListNode) {
	for n1 != nil && n2 != nil {
		t1, t2 := n1.Next, n2.Next
		n2.Next = t1
		n1.Next = n2
		n1, n2 = t1, t2
	}
}

func reverse(h *ListNode) *ListNode {
	var p *ListNode
	for h != nil {
		h, h.Next, p = h.Next, p, h
	}
	return p
}
