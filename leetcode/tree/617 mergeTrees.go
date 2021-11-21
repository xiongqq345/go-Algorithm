package tree

// 给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。
//
//你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点。
func mergeTrees(n1 *TreeNode, n2 *TreeNode) *TreeNode {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}

	return &TreeNode{
		Val:   n1.Val + n2.Val,
		Left:  mergeTrees(n1.Left, n2.Left),
		Right: mergeTrees(n1.Right, n2.Right),
	}
}
