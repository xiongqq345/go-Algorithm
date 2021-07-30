package dp

import "math"

// 假如有一排房子，共 n 个，每个房子可以被粉刷成 k 种颜色中的一种，你需要粉刷所有的房子并且使其相邻的两个房子颜色不能相同。
//
//当然，因为市场上不同颜色油漆的价格不同，所以房子粉刷成不同颜色的花费成本也是不同的。每个房子粉刷成不同颜色的花费是以一个 n x k 的矩阵来表示的。
//
//例如，costs[0][0] 表示第 0 号房子粉刷成 0 号颜色的成本花费；costs[1][2] 表示第 1 号房子粉刷成 2 号颜色的成本花费，以此类推。请你计算出粉刷完所有房子最少的花费成本。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/paint-house-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func minCostII(costs [][]int) int {
	n, k := len(costs), len(costs[0])
	for i := n - 2; i >= 0; i-- {
		for j := 0; j < k; j++ {
			costs[i][j] += minColor(costs[i+1], j)
		}
	}
	return minColor(costs[0], k+1)
}

// 找到当前房子的涂色方案中，不是 nextColor 方案的最小代价
func minColor(colors []int, nextColor int) int {
	// 找最小，初始值为最大
	ans := math.MaxInt64
	for i, cost := range colors {
		if i != nextColor {
			if ans > cost {
				ans = cost
			}
		}
	}
	return ans
}
