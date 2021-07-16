package backtracking

//给定一个字符串S，通过将字符串S中的每个字母转变大小写，我们可以获得一个新的字符串。返回所有可能得到的字符串集合。
func letterCasePermutation(s string) []string {
	n := len(s)
	var ans []string
	var helper func([]byte, int)
	helper = func(path []byte, pos int) {
		if pos == n {
			ans = append(ans, string(path))
			return
		}

		if s[pos] >= 'a' && s[pos] <= 'z' || s[pos] >= 'A' && s[pos] <= 'Z' {
			helper(append(path, s[pos]^32), pos+1)
		}
		helper(append(path, s[pos]), pos+1)
	}
	helper([]byte{}, 0)
	return ans
}
