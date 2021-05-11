package bt

// 请完成一个函数，输入一个二叉树，该函数输出它的镜像。
func invertTree(n *TreeNode) *TreeNode {
	if n == nil {
		return nil
	}
	n.Left, n.Right = n.Right, n.Left
	invertTree(n.Left)
	invertTree(n.Right)
	return n
}
