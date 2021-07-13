package coding

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 编写一个函数，检查输入的链表是否是回文的。
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	mid := midNode(head)
	tail := reverse(mid)
	p1, p2 := head, tail
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1, p2 = p1.Next, p2.Next
	}
	return true
}

func midNode(h *ListNode) *ListNode {
	p1, p2 := h, h.Next
	for p2 != nil && p2.Next != nil {
		p1, p2 = p1.Next, p2.Next.Next
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
