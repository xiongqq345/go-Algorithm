package linked_list

// 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
//
//你应当 保留 两个分区中每个节点的初始相对位置。
//
//输入：head = [1,4,3,2,5,2], x = 3
//输出：[1,2,2,4,3,5]

// [1,4,3,0,2,5,2] 3
// [1,0,2,2,4,3,5]

func partition(head *ListNode, x int) *ListNode {
	small, large := new(ListNode), new(ListNode)
	p1, p2 := small, large
	for head != nil {
		if head.Val < x {
			p1.Next, p1 = head, head
		} else {
			p2.Next, p2 = head, head
		}
		head = head.Next
	}
	p2.Next = nil
	p1.Next = large.Next
	return small.Next
}
