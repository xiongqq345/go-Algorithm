package linked_list

//输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 8 -> 0 -> 7
func addTwoNumbers3(l1 *ListNode, l2 *ListNode) *ListNode {
	s1, s2, carry := pushStack(l1), pushStack(l2), 0
	var newHead *ListNode
	for len(*s1) > 0 || len(*s2) > 0 {
		sum := popStack(s1) + popStack(s2) + carry
		carry, sum = sum/10, sum%10
		newHead = &ListNode{sum, newHead}
	}
	if carry > 0 {
		newHead = &ListNode{carry, newHead}
	}
	return newHead
}

func pushStack(head *ListNode) *[]int {
	var stack []int
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}
	return &stack
}

func popStack(stack *[]int) int {
	var res int
	if len(*stack) > 0 {
		res = (*stack)[len(*stack)-1]
		*stack = (*stack)[:len(*stack)-1]
	}
	return res
}
