package jz_offer

// 输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。
func pathSum(root *TreeNode, target int) [][]int {
	var ans [][]int
	var preorder func(*TreeNode, []int, int)
	preorder = func(node *TreeNode, path []int, left int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil && left == node.Val {
			ans = append(ans, append([]int{}, path...))
			return
		}
		preorder(node.Left, path, left-node.Val)
		preorder(node.Right, path, left-node.Val)
	}
	preorder(root, nil, target)
	return ans
}
