package _9_双周赛

import "math"

func countPaths(n int, roads [][]int) int {
	mod := 1e9 + 7
	const inf = math.MaxInt32
	dijkstra := make([][]int, n)
	for i := range dijkstra {
		dijkstra[i] = make([]int, n)
		for j := range dijkstra[i] {
			dijkstra[i][j] = inf
		}
	}
	for _, v := range roads {
		dijkstra[v[0]][v[1]] = v[2]
		dijkstra[v[1]][v[0]] = v[2]
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
