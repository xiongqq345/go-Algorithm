package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	var head, tail *ListNode
	for l1 != nil || l2 != nil {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum := v1 + v2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{
				sum,
				nil,
			}
			tail = head
		} else {
			tail.Next = &ListNode{sum, nil}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{carry, nil}
	}
	return head
}
