package _64_周赛

import "strconv"

//如果整数  x 满足：对于每个数位 d ，这个数位 恰好 在 x 中出现 d 次。那么整数 x 就是一个 数值平衡数 。
//
//给你一个整数 n ，请你返回 严格大于 n 的 最小数值平衡数 。
//
func nextBeautifulNumber(n int) int {
p:
	for i := n + 1; i < 1e9; i++ {
		str := strconv.Itoa(i)
		mp := make(map[int32]int32)
		for _, ch := range str {
			mp[ch-'0']++
		}
		for k, v := range mp {
			if k != v {
				continue p
			}
		}
		return i
	}
	return 0
}
