package string

import "strings"

// 给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。
//
//元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现。
func reverseVowels(s string) string {
	str := []byte(s)
	for i, j := 0, len(s)-1; i < j; {
		for i < len(s) && !strings.Contains("aeiouAEIOU", string(s[i])) {
			i++
		}
		for j > 0 && !strings.Contains("aeiouAEIOU", string(s[j])) {
			j--
		}
		if i < j {
			str[i], str[j] = str[j], str[i]
			i++
			j--
		}
	}
	return string(str)
}
