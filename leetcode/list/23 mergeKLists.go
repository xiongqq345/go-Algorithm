package list

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

//func merge(n1, n2 *ListNode) *ListNode {
//	dummy := new(ListNode)
//	t := dummy
//	for n1 != nil && n2 != nil {
//		if n1.Val < n2.Val {
//			t.Next, n1 = n1, n1.Next
//		} else {
//			t.Next, n2 = n2, n2.Next
//		}
//		t = t.Next
//	}
//	if n1 == nil {
//		t.Next = n2
//	} else {
//		t.Next = n1
//	}
//	return dummy.Next
//}
