package bt

func kthSmallest(root *TreeNode, k int) int {
	var vals []int
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		preorder(node.Left)
		vals = append(vals, node.Val)
		preorder(node.Right)
	}
	preorder(root)
	return vals[k-1]
}

func kthSmallest2(root *TreeNode, k int) int {
	var stack []*TreeNode
	var vals []int
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		vals = append(vals, node.Val)
		if len(vals) == k {
			return vals[k-1]
		}
		node = node.Right
	}
	return 0
}
