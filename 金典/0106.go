package coding

import (
	"strconv"
	"strings"
)

// 字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。比如，字符串aabcccccaaa会变为a2b1c5a3。若“压缩”后的字符串没有变短，则返回原先的字符串。你可以假设字符串中只包含大小写英文字母（a至z）。
func compressString(S string) string {
	var b strings.Builder
	var candidate rune
	var cnt int
	for _, v := range S {
		if v != candidate {
			if cnt > 0 {
				b.WriteRune(candidate)
				b.WriteString(strconv.Itoa(cnt))
			}
			candidate = v
			cnt = 0
		}
		cnt++
	}
	if cnt > 0 {
		b.WriteRune(candidate)
		b.WriteString(strconv.Itoa(cnt))
	}
	if b.Len() < len(S) {
		return b.String()
	}
	return S
}
