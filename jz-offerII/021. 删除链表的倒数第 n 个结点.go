package jz_offerII

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	p1, p2 := dummy, dummy
	for n > 0 {
		p2 = p2.Next
		n--
	}
	for p2.Next != nil {
		p1, p2 = p1.Next, p2.Next
	}
	p1.Next = p1.Next.Next
	return dummy.Next
}
