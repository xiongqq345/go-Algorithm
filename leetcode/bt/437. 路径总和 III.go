package bt

//给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
//
//路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
//

func pathSum(root *TreeNode, targetSum int) int {
	mp := make(map[int]int)
	mp[0] = 1
	var ans int
	var helper func(*TreeNode, int)
	helper = func(node *TreeNode, cur int) {
		if node == nil {
			return
		}
		cur += node.Val
		ans += mp[cur-targetSum]
		mp[cur]++
		helper(node.Left, cur)
		helper(node.Right, cur)
		mp[cur]--
	}
	helper(root, 0)
	return ans
}
