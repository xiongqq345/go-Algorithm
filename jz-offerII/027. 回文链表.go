package jz_offerII

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//给定一个链表的 头节点 head ，请判断其是否为回文链表。
//
//如果一个链表是回文，那么链表节点序列从前往后看和从后往前看是相同的。
func isPalindrome(head *ListNode) bool {
	mid := midNode(head)
	h2 := reverseList(mid.Next)
	mid.Next = nil
	p1, p2 := head, h2
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1, p2 = p1.Next, p2.Next
	}
	return true
}
