package dfs_bfs

func findCircleNum(isConnected [][]int) int {
	vis := make([]bool, len(isConnected))
	var cnt int
	var dfs func(int)
	dfs = func(city int) {
		vis[city] = true
		for j, v := range isConnected[city] {
			if v == 1 && !vis[j] {
				dfs(j)
			}
		}
	}
	for i := range isConnected {
		if !vis[i] {
			cnt++
			dfs(i)
		}
	}
	return cnt
}
