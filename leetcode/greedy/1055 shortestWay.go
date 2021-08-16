package greedy

// 对于任何字符串，我们可以通过删除其中一些字符（也可能不删除）来构造该字符串的子序列。
//
//给定源字符串 source 和目标字符串 target，找出源字符串中能通过串联形成目标字符串的子序列的最小数量。如果无法通过串联源字符串中的子序列来构造目标字符串，则返回 -1。
//
func shortestWay(source string, target string) int {
	var i int
	cnt := 1
	for _, v := range target {
		var f bool
		for j := i; j <= len(source); j++ {
			if f && j == i {
				return -1
			}
			if j == len(source) {
				j = 0
				cnt++
				f = true
			}
			if int32(source[j]) == v {
				i = j + 1
				break
			}
		}
	}
	return cnt
}
