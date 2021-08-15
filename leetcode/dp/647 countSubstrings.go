package dp

// 给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
// 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
func countSubstrings(s string) int {
	var cnt int
	for i := range s {
		cnt += palindrome(s, i, i) + palindrome(s, i, i+1)
	}
	return cnt
}

func palindrome(s string, i, j int) int {
	var cnt int
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
		cnt++
	}
	return cnt
}
