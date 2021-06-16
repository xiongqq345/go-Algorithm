package jz_offer

// 输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。
func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := height(root.Left), height(root.Right)
	if l == -1 || r == -1 || abs(l-r) > 1 {
		return -1
	}
	return max(l, r) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
