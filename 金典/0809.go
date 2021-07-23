package coding

// 括号。设计一种算法，打印n对括号的所有合法的（例如，开闭一一对应）组合。
//
//说明：解集不能包含重复的子集。
func generateParenthesis(n int) []string {
	var ans []string
	var helper func(string, int, int)
	helper = func(path string, l, r int) {
		if l > n || r > l {
			return
		}
		if len(path) == n*2 {
			ans = append(ans, path)
			return
		}
		helper(path+"(", l+1, r)
		helper(path+")", l, r+1)
	}
	helper("", 0, 0)
	return ans
}
