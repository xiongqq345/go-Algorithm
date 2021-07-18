package bt

// 给定一个二叉树，它的每个结点都存放着一个整数值。
//
//找出路径和等于给定数值的路径总数。
//
//路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
//
//二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数。
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
