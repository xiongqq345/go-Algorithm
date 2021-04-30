package bt

// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

func lowestCommonAncestor3(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q || root == nil {
		return root
	}
	left := lowestCommonAncestor3(root.Left, p, q)
	right := lowestCommonAncestor3(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}
