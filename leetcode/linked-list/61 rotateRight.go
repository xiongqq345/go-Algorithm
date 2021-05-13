package linked_list

// 给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
// [1,2,3,4] 2 [3,4,1,2]
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	p, n := head, 1
	for p.Next != nil {
		p = p.Next
		n++
	}
	p.Next = head
	p = head
	for i := 0; i < n-k%n-1; i++ {
		p = p.Next
	}
	res := p.Next
	p.Next = nil
	return res
}
