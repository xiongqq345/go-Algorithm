package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
	var ans int
	var inorder func(*TreeNode) int
	inorder = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := inorder(node.Left)
		r := inorder(node.Right)
		ans += abs(l - r)
		return l + r + node.Val
	}
	inorder(root)
	return ans
}
