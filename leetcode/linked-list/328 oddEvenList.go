package linked_list

//给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。
//输入: 1->2->3->4->5->NULL
//输出: 1->3->5->2->4->NULL
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	evenHead := head.Next
	o, e := head, head.Next
	for e != nil && e.Next != nil {
		o, o.Next = e.Next, e.Next
		e, e.Next = o.Next, o.Next
	}
	o.Next = evenHead
	return head
}
