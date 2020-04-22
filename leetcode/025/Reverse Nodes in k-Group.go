package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

//反转 a 到 b 之间的结点
func reverse(head *ListNode, b *ListNode) *ListNode {
	var preNode *ListNode
	preNode = nil
	//  后一个节点
	nextNode := new(ListNode)
	nextNode = nil
	for head != b {
		nextNode = head.Next
		head.Next = preNode
		preNode = head
		head = nextNode
	}
	return preNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}

	newhead := reverse(a, b)
	a.Next = reverseKGroup(b, k)
	return newhead
}
