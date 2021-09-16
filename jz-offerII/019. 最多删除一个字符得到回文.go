package jz_offerII

//给定一个非空字符串 s，请判断如果 最多 从字符串中删除一个字符能否得到一个回文字符串。
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
