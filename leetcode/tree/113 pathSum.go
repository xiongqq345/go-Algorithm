package tree

// 给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	var dfs func(*TreeNode, int, []int)
	dfs = func(node *TreeNode, left int, path []int) {
		if node == nil {
			return
		}

		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil && node.Val == left {
			res = append(res, append([]int{}, path...))
		}
		dfs(node.Left, left-node.Val, path)
		dfs(node.Right, left-node.Val, path)

	}
	dfs(root, targetSum, []int{})
	return res
}
