package linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	p1, p2 := head, head
	for p2 != nil && p2.Next != nil {
		p1, p2 = p1.Next, p2.Next.Next
	}
	return p1
}
