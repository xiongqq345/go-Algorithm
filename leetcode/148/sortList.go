package _48

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := midNode(head)
	l := sortList(mid.Next)
	mid.Next = nil
	r := sortList(head)
	return merge(l, r)
}

func midNode(head *ListNode) *ListNode {
	p1, p2 := head, head.Next
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return p1
}

func merge(l, r *ListNode) *ListNode {
	dummy := new(ListNode)
	t := dummy

	for l != nil && r != nil {
		if l.Val < r.Val {
			t.Next = l
			l = l.Next
		} else {
			t.Next = r
			r = r.Next
		}
		t = t.Next
	}
	if l != nil {
		t.Next = l
	} else {
		t.Next = r
	}

	return dummy.Next
}
