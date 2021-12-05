package _70_周赛

//给你一个链表的头节点 head 。删除 链表的 中间节点 ，并返回修改后的链表的头节点 head 。
//
//长度为 n 链表的中间节点是从头数起第 ⌊n / 2⌋ 个节点（下标从 0 开始），其中 ⌊x⌋ 表示小于或等于 x 的最大整数。
//
//对于 n = 1、2、3、4 和 5 的情况，中间节点的下标分别是 0、1、1、2 和 2 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/delete-the-middle-node-of-a-linked-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev, mid := midNode2(dummy), midNode(head)
	prev.Next = mid.Next
	return dummy.Next
}

func midNode(head *ListNode) *ListNode {
	p1, p2 := head, head
	for p2 != nil && p2.Next != nil {
		p1, p2 = p1.Next, p2.Next.Next
	}
	return p1
}

func midNode2(head *ListNode) *ListNode {
	p1, p2 := head, head.Next
	for p2 != nil && p2.Next != nil {
		p1, p2 = p1.Next, p2.Next.Next
	}
	return p1
}
