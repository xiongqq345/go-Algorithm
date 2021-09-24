package jz_offerII

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//给定一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
//
//每条从根节点到叶节点的路径都代表一个数字：
//
//例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
//计算从根节点到叶节点生成的 所有数字之和 。
//
//叶节点 是指没有子节点的节点。
//
func sumNumbers(root *TreeNode) int {
	var sum int
	var preorder func(*TreeNode, int)
	preorder = func(node *TreeNode, path int) {
		if node == nil {
			return
		}
		path = 10*path + node.Val
		if node.Left == nil && node.Right == nil {
			sum += path
			return
		}
		preorder(node.Left, path)
		preorder(node.Right, path)
	}
	preorder(root, 0)
	return sum
}
