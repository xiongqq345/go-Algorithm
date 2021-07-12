package coding

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。
func kthToLast(head *ListNode, k int) int {
	p1, p2 := head, head
	for i := 0; i < k; i++ {
		p2 = p2.Next
	}
	for p2 != nil {
		p1, p2 = p1.Next, p2.Next
	}
	return p1.Val
}
