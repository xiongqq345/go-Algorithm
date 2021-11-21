package tree

//给定一个二叉树
//
//struct Node {
//  int val;
//  Node *left;
//  Node *right;
//  Node *next;
//}
//填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
//
//初始状态下，所有 next 指针都被设置为 NULL。

func connect2(root *Node) *Node {
	start := root
	for start != nil {
		var nextStart, last *Node
		handle := func(node *Node) {
			if node == nil {
				return
			}
			if nextStart == nil {
				nextStart = node
			}
			if last != nil {
				last.Next = node
			}
			last = node
		}

		for p := start; p != nil; p = p.Next {
			handle(p.Left)
			handle(p.Right)
		}
		start = nextStart
	}
	return root
}
