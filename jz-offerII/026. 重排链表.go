package jz_offerII

//给定一个单链表 L 的头节点 head ，单链表 L 表示为：
//
// L0 → L1 → … → Ln-1 → Ln
//请将其重新排列后变为：
//
//L0 → Ln → L1 → Ln-1 → L2 → Ln-2 → …
//
//不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/LGjMqU
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(h *ListNode) {
	mid := midNode(h)
	h2 := reverseList(mid.Next)
	mid.Next = nil
	p1, p2 := h, h2
	for p1 != nil && p2 != nil {
		n1, n2 := p1.Next, p2.Next
		p1.Next = p2
		p2.Next = n1
		p1, p2 = n1, n2
	}
}
