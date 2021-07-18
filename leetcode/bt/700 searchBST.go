package bt

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
	node := root
	for node != nil {
		if val < node.Val {
			node = node.Left
		} else if val > node.Val {
			node = node.Right
		} else {
			return node
		}
	}
	return nil
}
