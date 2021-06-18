package backtracking

// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
func generateParenthesis(n int) []string {
	var ans []string
	var dfs func(path string, l, r int)
	dfs = func(path string, l, r int) {
		if l > n || r > l {
			return
		}
		if len(path) >= n*2 {
			ans = append(ans, path)
			return
		}
		dfs(path+"(", l+1, r)
		dfs(path+")", l, r+1)
	}
	dfs("", 0, 0)
	return ans
}
