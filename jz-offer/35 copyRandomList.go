package jz_offer

//请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。
func copyRandomList(head *Node) *Node {
	mp := make(map[*Node]*Node)
	dummy := new(Node)
	p, cur := dummy, head
	for cur != nil {
		node := &Node{Val: cur.Val}
		p, p.Next, mp[cur] = node, node, node
		cur = cur.Next
	}

	p, cur = dummy.Next, head
	for cur != nil {
		p.Random = mp[cur.Random]
		p, cur = p.Next, cur.Next
	}

	return dummy.Next
}
