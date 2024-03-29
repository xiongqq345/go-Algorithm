package array

import "math"

// 在一个小镇里，按从 1 到 n 为 n 个人进行编号。传言称，这些人中有一个是小镇上的秘密法官。
//
//如果小镇的法官真的存在，那么：
//
//小镇的法官不相信任何人。
//每个人（除了小镇法官外）都信任小镇的法官。
//只有一个人同时满足条件 1 和条件 2 。
//给定数组 trust，该数组由信任对 trust[i] = [a, b] 组成，表示编号为 a 的人信任编号为 b 的人。
//
//如果小镇存在秘密法官并且可以确定他的身份，请返回该法官的编号。否则，返回 -1。
func findJudge(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}

	set := make(map[int]bool)
	mp := make(map[int]int)
	for _, v := range trust {
		set[v[0]] = true
		mp[v[1]]++
	}

	maxTrust := math.MinInt32
	var judge int
	for k, v := range mp {
		if maxTrust < v {
			maxTrust = v
			judge = k
		}
	}
	if maxTrust == n-1 && !set[judge] {
		return judge
	}

	return -1
}
