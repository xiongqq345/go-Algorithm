package jz_offerII

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
func findBottomLeftValue(root *TreeNode) int {
	var d int
	ans := root.Val
	var preorder func(*TreeNode, int)
	preorder = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth > d {
			d = depth
			ans = node.Val
		}
		preorder(node.Left, depth+1)
		preorder(node.Right, depth+1)
	}
	preorder(root, 0)
	return ans
}
