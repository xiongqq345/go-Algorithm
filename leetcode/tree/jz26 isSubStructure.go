package tree

// 树的子结构
func isSubStructure(a *TreeNode, b *TreeNode) bool {
	if a == nil || b == nil {
		return false
	}
	if recur(a, b) {
		return true
	}
	return isSubStructure(a.Left, b) || isSubStructure(a.Right, b)
}

func recur(a *TreeNode, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil || a.Val != b.Val {
		return false
	}
	return recur(a.Left, b.Left) && recur(a.Right, b.Right)
}
