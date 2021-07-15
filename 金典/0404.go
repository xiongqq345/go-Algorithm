package coding

// 实现一个函数，检查二叉树是否平衡。在这个问题中，平衡树的定义如下：任意一个节点，其两棵子树的高度差不超过 1

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	l, r := height(node.Left), height(node.Right)
	if l == -1 || r == -1 || abs(l-r) > 1 {
		return -1
	}
	return max(l, r) + 1
}
