package coding

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 给定一个链表，如果它是有环链表，实现一个算法返回环路的开头节点。
//
//如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
//
func detectCycle(head *ListNode) *ListNode {
	p, p1, p2 := head, head, head
	for p2 != nil && p2.Next != nil {
		p1, p2 = p1.Next, p2.Next.Next
		if p1 == p2 {
			break
		}
	}

	if p2 == nil || p2.Next == nil {
		return nil
	}
	for p != p1 {
		p, p1 = p.Next, p1.Next
	}
	return p
}
