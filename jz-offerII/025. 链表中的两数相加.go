package jz_offerII

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//给定两个 非空链表 l1和 l2 来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
//
//可以假设除了数字 0 之外，这两个数字都不会以零开头。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/lMSNwu
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverseList(l1)
	l2 = reverseList(l2)
	var sum, carry int
	dummy := new(ListNode)
	p := dummy
	for l1 != nil || l2 != nil || carry > 0 {
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum = v1 + v2 + carry
		sum, carry = sum%10, sum/10
		p.Next = &ListNode{Val: sum}
		p = p.Next
	}

	return reverseList(dummy.Next)
}
