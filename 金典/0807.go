package coding

// 无重复字符串的排列组合。编写一种方法，计算某字符串的所有排列组合，字符串每个字符均不相同。
func permutation(s string) []string {
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
			if !vis[i] {
				vis[i] = true
				path = append(path, s[i])
				helper(pos + 1)
				path = path[:len(path)-1]
				vis[i] = false
			}
		}
	}
	helper(0)
	return ans
}
