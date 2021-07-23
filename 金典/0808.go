package coding

import "sort"

// 有重复字符串的排列组合。编写一种方法，计算某字符串的所有排列组合。
func permutation(s string) []string {
	str := []byte(s)
	sort.Slice(str, func(i, j int) bool {
		return str[i] < str[j]
	})
	n := len(s)
	var path []byte
	var ans []string
	vis := make([]bool, len(s))
	var helper func(int)
	helper = func(pos int) {
		if pos == n {
			ans = append(ans, string(path))
			return
		}

		for i := 0; i < n; i++ {
			if i > 0 && str[i] == str[i-1] && vis[i-1] {
				continue
			}
			if !vis[i] {
				vis[i] = true
				path = append(path, str[i])
				helper(pos + 1)
				path = path[:len(path)-1]
				vis[i] = false
			}
		}
	}
	helper(0)
	return ans
}
