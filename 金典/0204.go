package coding

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 编写程序以 x 为基准分割链表，使得所有小于 x 的节点排在大于或等于 x 的节点之前。如果链表中包含 x，x 只需出现在小于 x 的元素之后(如下所示)。分割元素 x 只需处于“右半部分”即可，其不需要被置于左右两部分之间。
//
func partition(head *ListNode, x int) *ListNode {
	dummy1, dummy2 := new(ListNode), new(ListNode)
	p, p1, p2 := head, dummy1, dummy2
	for p != nil {
		if p.Val < x {
			p1.Next = p
			p1 = p1.Next
		} else {
			p2.Next = p
			p2 = p2.Next
		}
		p = p.Next
	}
	p2.Next = nil
	p1.Next = dummy2.Next
	return dummy1.Next
}
