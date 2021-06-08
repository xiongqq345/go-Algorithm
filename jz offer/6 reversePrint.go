package jz_offer

// 输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	return append(reversePrint(head.Next), head.Val)
}
