package _4__双周赛

import "sort"

// 给你一个二维整数数组 ranges 和两个整数 left 和 right 。每个 ranges[i] = [starti, endi] 表示一个从 starti 到 endi 的 闭区间 。
//
//如果闭区间 [left, right] 内每个整数都被 ranges 中 至少一个 区间覆盖，那么请你返回 true ，否则返回 false 。
//
//已知区间 ranges[i] = [starti, endi] ，如果整数 x 满足 starti <= x <= endi ，那么我们称整数x 被覆盖了。
func isCovered(ranges [][]int, left int, right int) bool {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	var p int
	for p < len(ranges) && left <= right {
		switch {
		case left >= ranges[p][0] && left <= ranges[p][1]:
			left = ranges[p][1]
		case left < ranges[p][0]:
			return false
		case left > ranges[p][1]:
			p++
		}
	}

	return left > right
}
