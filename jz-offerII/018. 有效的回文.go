package jz_offerII

import "strings"

//给定一个字符串 s ，验证 s 是否是 回文串 ，只考虑字母和数字字符，可以忽略字母的大小写。
func isPalindrome(s string) bool {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		if isalnum(s[i]) {
			b.WriteByte(s[i])
		}
	}

	s = b.String()
	n := len(s)
	s = strings.ToLower(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
