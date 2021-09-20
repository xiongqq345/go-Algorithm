package jz_offerII

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(h *ListNode) *ListNode {
	p, p1, p2 := h, h, h
	for p2 != nil && p2.Next != nil {
		p1, p2 = p1.Next, p2.Next.Next
		if p1 == p2 {
			for p != p1 {
				p, p1 = p.Next, p1.Next
			}
			return p
		}
	}
	return nil
}
