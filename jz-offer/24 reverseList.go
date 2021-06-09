package jz_offer

func reverseList(h *ListNode) *ListNode {
	var p *ListNode
	for h != nil {
		h, h.Next, p = h.Next, p, h
	}
	return p
}
