package jz_offer

// 给定一棵二叉搜索树，请找出其中第k大的节点。
func kthLargest(root *TreeNode, k int) int {
	var vals []int
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Right)
		vals = append(vals, node.Val)
		if len(vals) == k {
			return
		}
		inorder(node.Left)
	}
	inorder(root)
	return vals[k-1]
}
