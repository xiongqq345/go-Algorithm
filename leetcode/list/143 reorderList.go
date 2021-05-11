package list

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

func mergeList(h1, h2 *ListNode) {
	for h1 != nil && h2 != nil {
		t1, t2 := h1.Next, h2.Next
		h2.Next = t1
		h1.Next = h2
		h1, h2 = t1, t2
	}
}

func midNode(h *ListNode) *ListNode {
	p1, p2 := h, h.Next
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return p1
}

func reverse(h *ListNode) *ListNode {
	var p *ListNode
	for h != nil {
		h, h.Next, p = h.Next, p, h
	}
	return p
}
