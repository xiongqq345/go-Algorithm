package string

import "strings"

// 给定一个字符串 s 和一个整数 k，从字符串开头算起，每 2k 个字符反转前 k 个字符。
//
//如果剩余字符少于 k 个，则将剩余字符全部反转。
//如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。

func reverseStr(s string, k int) string {
	str := []byte(s)
	var b strings.Builder
	for i := 0; i < len(s); i += 2 * k {
		if len(s)-i >= k {
			b.Write(rev(str[i : i+k]))
			if i+2*k > len(s) {
				b.Write(str[i+k:])
			} else {
				b.Write(str[i+k : i+2*k])
			}
		} else {
			b.Write(rev(str[i:]))
			break
		}
	}
	return b.String()
}

func rev(s []byte) []byte {
	n := len(s)
	for i := range s[:n/2] {
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}
	return s
}
