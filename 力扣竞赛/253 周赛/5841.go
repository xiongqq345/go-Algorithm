package _53_周赛

import "sort"

// 你打算构建一些障碍赛跑路线。给你一个 下标从 0 开始 的整数数组 obstacles ，数组长度为 n ，其中 obstacles[i] 表示第 i 个障碍的高度。
//
//对于每个介于 0 和 n - 1 之间（包含 0 和 n - 1）的下标  i ，在满足下述条件的前提下，请你找出 obstacles 能构成的最长障碍路线的长度：
//
//你可以选择下标介于 0 到 i 之间（包含 0 和 i）的任意个障碍。
//在这条路线中，必须包含第 i 个障碍。
//你必须按障碍在 obstacles 中的 出现顺序 布置这些障碍。
//除第一个障碍外，路线中每个障碍的高度都必须和前一个障碍 相同 或者 更高 。
//返回长度为 n 的答案数组 ans ，其中 ans[i] 是上面所述的下标 i 对应的最长障碍赛跑路线的长度。
func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	var ans, dp []int
	for _, v := range obstacles {
		p := sort.SearchInts(dp, v+1)
		ans = append(ans, p+1)
		if p < len(dp) {
			dp[p] = v
		} else {
			dp = append(dp, v)
		}
	}
	return ans
}
