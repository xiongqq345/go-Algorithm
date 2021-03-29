package _226

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(n *TreeNode) *TreeNode {
	if n == nil {
		return nil
	}
	n.Left, n.Right = n.Right, n.Left
	invertTree(n.Left)
	invertTree(n.Right)
	return n
}
