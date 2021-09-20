package jz_offerII

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(h *ListNode) (p *ListNode) {
	for h != nil {
		h, h.Next, p = h.Next, p, h
	}
	return
}
