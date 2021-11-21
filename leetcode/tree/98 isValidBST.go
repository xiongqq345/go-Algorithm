package tree

import "math"

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}

	pre := math.MinInt64
	var inorder func(*TreeNode) bool
	inorder = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if !inorder(node.Left) {
			return false
		}
		if node.Val <= pre {
			return false
		}
		pre = node.Val
		return inorder(node.Right)
	}
	return inorder(root)
}
