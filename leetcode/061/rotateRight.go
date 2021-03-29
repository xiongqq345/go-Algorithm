package _61

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	p := head
	n := 1
	for p.Next != nil {
		p = p.Next
		n++
	}
	p.Next = head
	p = head
	for i := 0; i < n-1-k%n; i++ {
		p = p.Next
	}
	res := p.Next
	p.Next = nil
	return res
}
