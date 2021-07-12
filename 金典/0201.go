package coding

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。
func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	set := make(map[int]bool)
	set[head.Val] = true
	p := head
	for p.Next != nil {
		if set[p.Next.Val] {
			p.Next = p.Next.Next
		} else {
			set[p.Next.Val] = true
			p = p.Next
		}
	}
	return head
}
