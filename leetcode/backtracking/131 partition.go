package backtracking

// 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
//
//回文串 是正着读和反着读都一样的字符串。
func partition(s string) [][]string {
	var ans [][]string
	n := len(s)
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}

	var splits []string
	var helper func(int)
	helper = func(i int) {
		if i == n {
			ans = append(ans, append([]string{}, splits...))
			return
		}

		for j := i; j < n; j++ {
			if f[i][j] {
				splits = append(splits, s[i:j+1])
				helper(j + 1)
				splits = splits[:len(splits)-1]
			}
		}
	}
	helper(0)
	return ans
}
