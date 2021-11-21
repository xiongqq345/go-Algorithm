package tree

import "strconv"

// 给定一个二叉树，返回所有从根节点到叶子节点的路径。
func binaryTreePaths(root *TreeNode) []string {
	var res []string
	var preorder func(*TreeNode, string)
	preorder = func(node *TreeNode, path string) {
		if node == nil {
			return
		}
		path += strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			res = append(res, path)
		} else {
			path += "->"
			preorder(node.Left, path)
			preorder(node.Right, path)
		}
	}
	preorder(root, "")
	return res
}
