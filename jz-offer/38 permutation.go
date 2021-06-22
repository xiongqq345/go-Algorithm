package jz_offer

import "sort"

// 输入一个字符串，打印出该字符串中字符的所有排列。
//你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。
func permutation(s string) []string {
	var ans []string
	var path []byte
	str := []byte(s)
	sort.Slice(str, func(i, j int) bool {
		return str[i] < str[j]
	})
	used := make([]bool, len(s))
	var helper func()
	helper = func() {
		if len(path) == len(s) {
			ans = append(ans, string(path))
			return
		}
		for i := range str {
			if i > 0 && str[i] == str[i-1] && !used[i-1] {
				continue
			}
			if !used[i] {
				path = append(path, str[i])
				used[i] = true
				helper()
				used[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	helper()
	return ans
}
