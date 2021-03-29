package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(n1 *TreeNode, n2 *TreeNode) *TreeNode {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	n1.Val += n2.Val
	n1.Left = mergeTrees(n1.Left, n2.Left)
	n1.Right = mergeTrees(n1.Right, n2.Right)
	return n1
}
