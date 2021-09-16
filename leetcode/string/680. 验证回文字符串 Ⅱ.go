package string

//给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。
func validPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			if isPalindrome(s[:i]+s[i+1:]) || isPalindrome(s[:n-1-i]+s[n-i:]) {
				return true
			}
			return false
		}
	}
	return true
}

func isPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}
