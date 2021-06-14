package jz_offer

//输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var depth int
	q := []*TreeNode{root}
	for len(q) > 0 {
		var nq []*TreeNode
		for _, node := range q {
			if node.Left != nil {
				nq = append(nq, node.Left)
			}
			if node.Right != nil {
				nq = append(nq, node.Right)
			}
		}
		q = nq
		depth++
	}
	return depth
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
