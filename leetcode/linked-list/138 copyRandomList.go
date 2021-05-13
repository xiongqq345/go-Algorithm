package linked_list

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//构造这个链表的 深拷贝
func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)
	dummy := new(Node)
	p, cur := dummy, head
	for cur != nil {
		node := &Node{Val: cur.Val}
		m[cur] = node
		p.Next = node
		p = node
		cur = cur.Next
	}

	cur, p = head, dummy.Next
	for cur != nil {
		p.Random = m[cur.Random]
		cur, p = cur.Next, p.Next
	}
	return dummy.Next
}
