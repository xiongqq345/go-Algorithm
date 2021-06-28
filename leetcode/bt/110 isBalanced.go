package bt

// 给定一个二叉树，判断它是否是高度平衡的二叉树。
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if abs(height(root.Left)-height(root.Right)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left)+1, height(root.Right)+1)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func isBalanced2(root *TreeNode) bool {
	return recur(root) >= 0
}

func recur(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := recur(root.Left)
	if l == -1 {
		return -1
	}
	r := recur(root.Right)
	if r == -1 {
		return -1
	}
	if abs(l-r) > 1 {
		return -1
	}
	return max(l, r) + 1
}
