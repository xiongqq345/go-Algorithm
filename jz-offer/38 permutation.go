package jz_offer

import "sort"

// 输入一个字符串，打印出该字符串中字符的所有排列。
//你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。
func permutation(s string) []string {
	var ans []string
	str := []byte(s)
	sort.Slice(str, func(i, j int) bool {
		return str[i] < str[j]
	})
	used := make([]bool, len(s))
	var helper func([]byte)
	helper = func(path []byte) {
		if len(path) == len(s) {
			ans = append(ans, string(path))
			return
		}
		for i := range str {
			if i > 0 && str[i] == str[i-1] && !used[i-1] {
				continue
			}
			if !used[i] {
				used[i] = true
				helper(append(path, str[i]))
				used[i] = false
			}
		}
	}
	helper(nil)
	return ans
}
