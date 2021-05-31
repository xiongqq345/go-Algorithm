package linked_list

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	var preorder func(*Node, *[]int)
	preorder = func(node *Node, list *[]int) {
		if node == nil {
			return
		}
		*list = append(*list, node.Val)
		preorder(node.Child, list)
		preorder(node.Next, list)
	}
	var list []int
	preorder(root, &list)
	dummy := &Node{}
	cur := dummy
	for _, v := range list {
		node := &Node{
			Val:  v,
			Prev: cur,
		}
		cur, cur.Next = node, node
	}
	dummy.Next.Prev = nil
	return dummy.Next
}

func flatten2(root *Node) *Node {
	if root == nil {
		return nil
	}
	var cur, next *Node = root, nil
	for cur != nil {
		next = cur.Next
		if cur.Child != nil {
			child := cur.Child
			cur.Child, cur.Next, child.Prev = nil, cur.Child, cur
			for child.Next != nil {
				child = child.Next
			}
			child.Next = next
			if next != nil {
				next.Prev = child
			}
		}
		cur = cur.Next
	}
	return root
}
