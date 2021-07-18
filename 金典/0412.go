package coding

// 给定一棵二叉树，其中每个节点都含有一个整数数值(该值或正或负)。设计一个算法，打印节点数值总和等于某个给定值的所有路径的数量。注意，路径不一定非得从二叉树的根节点或叶节点开始或结束，但是其方向必须向下(只能从父节点指向子节点方向)。
//
var mp map[int]int

func pathSum(root *TreeNode, targetSum int) int {
	mp = make(map[int]int)
	mp[0] = 1
	return dfs(root, targetSum, 0)
}

func dfs(node *TreeNode, target, sum int) int {
	if node == nil {
		return 0
	}
	sum += node.Val
	result := mp[sum-target]
	mp[sum]++
	result += dfs(node.Left, target, sum) + dfs(node.Right, target, sum)
	mp[sum]--
	return result
}
