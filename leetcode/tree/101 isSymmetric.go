package tree

func isSymmetric(root *TreeNode) bool {
	return compare(root, root)
}

func compare(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && compare(p.Left, q.Right) && compare(p.Right, q.Left)
}
