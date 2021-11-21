package tree

// 给定一个整数 n，生成所有由 1 ... n 为节点所组成的 二叉搜索树 。
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return dfs(1, n)
}

func dfs(l, r int) []*TreeNode {
	if l > r {
		return []*TreeNode{nil}
	}
	var allTrees []*TreeNode
	for i := l; i <= r; i++ {
		lts := dfs(l, i-1)
		rts := dfs(i+1, r)
		for _, lt := range lts {
			for _, rt := range rts {
				allTrees = append(allTrees, &TreeNode{
					Val:   i,
					Left:  lt,
					Right: rt,
				})
			}
		}
	}
	return allTrees
}
