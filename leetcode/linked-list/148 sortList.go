package linked_list

// 给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := midNode(head)
	l1 := sortList(mid.Next)
	mid.Next = nil
	l2 := sortList(head)
	return merge(l1, l2)
}

func midNode(head *ListNode) *ListNode {
	p1, p2 := head, head.Next
	for p2 != nil && p2.Next != nil {
		p1, p2 = p1.Next, p2.Next.Next
	}
	return p1
}

func merge(n1, n2 *ListNode) *ListNode {
	dummy := new(ListNode)
	t := dummy
	for n1 != nil && n2 != nil {
		if n1.Val < n2.Val {
			t.Next, n1 = n1, n1.Next
		} else {
			t.Next, n2 = n2, n2.Next
		}
		t = t.Next
	}
	if n1 == nil {
		t.Next = n2
	} else {
		t.Next = n1
	}
	return dummy.Next
}
