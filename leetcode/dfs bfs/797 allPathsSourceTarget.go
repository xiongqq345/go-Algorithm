package dfs_bfs

// 给一个有 n 个结点的有向无环图，找到所有从 0 到 n-1 的路径并输出（不要求按顺序）
//
//二维数组的第 i 个数组中的单元都表示有向图中 i 号结点所能到达的下一些结点（译者注：有向图是有方向的，即规定了 a→b 你就不能从 b→a ）空就是没有下一个结点了。

func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	var ans [][]int
	var helper func(int, []int)
	helper = func(pos int, path []int) {
		if pos == n-1 {
			ans = append(ans, append([]int(nil), path...))
			return
		}

		for i := range graph[pos] {
			helper(graph[pos][i], append(path, graph[pos][i]))
		}
	}
	helper(0, []int{0})
	return ans
}
