package bt

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//给你二叉树的根结点 root ，此外树的每个结点的值要么是 0 ，要么是 1 。
//
//返回移除了所有不包含 1 的子树的原二叉树。
//
//节点 node 的子树为 node 本身加上所有 node 的后代。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-tree-pruning
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}
	return root
}
