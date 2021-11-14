package _67_周赛
//给你一个链表的头节点 head 。
//
//链表中的节点 按顺序 划分成若干 非空 组，这些非空组的长度构成一个自然数序列（1, 2, 3, 4, ...）。一个组的 长度 就是组中分配到的节点数目。换句话说：
//
//节点 1 分配给第一组
//节点 2 和 3 分配给第二组
//节点 4、5 和 6 分配给第三组，以此类推
//注意，最后一组的长度可能小于或者等于 1 + 倒数第二组的长度 。
//
//反转 每个 偶数 长度组中的节点，并返回修改后链表的头节点 head 。

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	return reverseKGroup(head, 1, false)
}

func reverseKGroup(head *ListNode, k int, even bool) *ListNode {
	p := head
	var prev *ListNode
	for i := 0; i < k; i++ {
		if p == nil {
			if i > 0 && i%2 == 0 {
				res := reverseList(head)
				head.Next = nil
				return res
			}
			return head
		}
		prev = p
		p = p.Next
	}
	if even {
		res := reverseToNode(head, p)
		head.Next = reverseKGroup(p, k+1, !even)
		return res
	}
	prev.Next = reverseKGroup(p, k+1, !even)
	return head
}

func reverseList(h *ListNode) *ListNode {
	var p *ListNode
	for h != nil {
		h, h.Next, p = h.Next, p, h
	}
	return p
}

func reverseToNode(h, node *ListNode) *ListNode {
	var p *ListNode
	for h != node {
		h, h.Next, p = h.Next, p, h
	}
	return p
}
