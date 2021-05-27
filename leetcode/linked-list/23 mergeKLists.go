package linked_list

// 给你一个链表数组，每个链表都已经按升序排列。
//
//请你将所有链表合并到一个升序链表中，返回合并后的链表。

func mergeKLists(lists []*ListNode) *ListNode {
	var res *ListNode
	for _, list := range lists {
		res = merge(res, list)
	}
	return res
}
