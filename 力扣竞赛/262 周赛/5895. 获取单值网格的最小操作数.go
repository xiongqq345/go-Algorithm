package _62_周赛

import "sort"

//给你一个大小为 m x n 的二维整数网格 grid 和一个整数 x 。每一次操作，你可以对 grid 中的任一元素 加 x 或 减 x 。
//
//单值网格 是全部元素都相等的网格。
//
//返回使网格化为单值网格所需的 最小 操作数。如果不能，返回 -1 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/minimum-operations-to-make-a-uni-value-grid
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func minOperations(grid [][]int, x int) int {
	var arr []int
	for i := range grid {
		arr = append(arr, grid[i]...)
	}
	sort.Ints(arr)
	n := len(arr)
	if n == 1 {
		return 0
	}
	var ans int
	for _, v := range arr {
		if (v-arr[0])%x > 0 {
			return -1
		}
		ans += abs(v-arr[n/2]) / x
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}
