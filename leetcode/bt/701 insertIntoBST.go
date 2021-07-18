package bt

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 给定二叉搜索树（BST）的根节点和要插入树中的值，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。 输入数据 保证 ，新值和原始二叉搜索树中的任意节点值都不同。
//
//注意，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。 你可以返回 任意有效的结果 。
//
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	node := root
	for node != nil {
		if val < node.Val {
			if node.Left != nil {
				node = node.Left
			} else {
				node.Left = &TreeNode{Val: val}
				break
			}
		} else {
			if node.Right != nil {
				node = node.Right
			} else {
				node.Right = &TreeNode{Val: val}
				break
			}
		}
	}
	return root
}
