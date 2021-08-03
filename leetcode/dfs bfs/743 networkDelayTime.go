package dfs_bfs

import "math"

// 有 n 个网络节点，标记为 1 到 n。
//
//给你一个列表 times，表示信号经过 有向 边的传递时间。 times[i] = (ui, vi, wi)，其中 ui 是源节点，vi 是目标节点， wi 是一个信号从源节点传递到目标节点的时间。
//
//现在，从某个节点 K 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1 。
//
func networkDelayTime(times [][]int, n int, k int) int {
	const inf = math.MaxInt32
	dijkstra := make([][]int, n)
	for i := range dijkstra {
		dijkstra[i] = make([]int, n)
		for j := range dijkstra[i] {
			dijkstra[i][j] = inf
		}
	}
	for _, v := range times {
		dijkstra[v[0]-1][v[1]-1] = v[2]
	}
	vis := make([]bool, n)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[k-1] = 0
	for i := 0; i < n; i++ {
		minDist := inf
		var minIndex int
		for j, v := range dist {
			if !vis[j] && v < minDist {
				minDist = v
				minIndex = j
			}
		}
		vis[minIndex] = true
		for j, v := range dijkstra[minIndex] {
			dist[j] = min(dist[j], v+minDist)
		}
	}

	var ans int
	for _, v := range dist {
		if v == inf {
			return -1
		}
		ans = max(ans, v)
	}
	return ans
}
