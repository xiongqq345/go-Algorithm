package greedy

import "strings"

// 给定一个以字符串表示的非负整数 num，移除这个数中的 k 位数字，使得剩下的数字最小。
func removeKdigits(num string, k int) string {
	var st []rune
	for _, n := range num {
		for len(st) > 0 && st[len(st)-1] > n && k > 0 {
			st = st[:len(st)-1]
			k--
		}
		st = append(st, n)
	}
	st = st[:len(st)-k]
	ans := strings.TrimLeft(string(st), "0")
	if ans == "" {
		return "0"
	}
	return ans
}
