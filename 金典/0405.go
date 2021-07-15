package coding

import "math"

// 实现一个函数，检查一棵二叉树是否为二叉搜索树。
//
//示例 1:
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
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
