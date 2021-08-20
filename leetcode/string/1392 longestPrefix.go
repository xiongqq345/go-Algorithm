package string

// 「快乐前缀」是在原字符串中既是 非空 前缀也是后缀（不包括原字符串自身）的字符串。
//
//给你一个字符串 s，请你返回它的 最长快乐前缀。
//
//如果不存在满足题意的前缀，则返回一个空字符串。
//
func longestPrefix(s string) string {
	for idx := len(s) - 1; idx > 0; idx-- {
		if s[:idx] == s[len(s)-idx:] {
			return s[:idx]
		}
	}
	return ""
}
