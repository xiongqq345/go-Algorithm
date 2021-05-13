package linked_list

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)
	dummy := new(Node)

	p, cur := dummy, head
	for cur != nil {
		node := &Node{Val: cur.Val}
		p, p.Next, m[cur] = node, node, node
		cur = cur.Next
	}

	p, cur = dummy.Next, head
	for cur != nil {
		p.Random = m[cur.Random]
		p, cur = p.Next, cur.Next
	}
	return dummy.Next
}
