package string

import (
	"strings"
)

//有一个密钥字符串 S ，只包含字母，数字以及 '-'（破折号）。其中， N 个 '-' 将字符串分成了 N+1 组。
//
//给你一个数字 K，请你重新格式化字符串，使每个分组恰好包含 K 个字符。特别地，第一个分组包含的字符个数必须小于等于 K，但至少要包含 1 个字符。两个分组之间需要用 '-'（破折号）隔开，并且将所有的小写字母转换为大写字母。
//
//给定非空字符串 S 和数字 K，按照上面描述的规则进行格式化。
//
func licenseKeyFormatting(s string, k int) string {
	arr := strings.Split(s, "-")
	s = strings.Join(arr, "")
	s = strings.ToUpper(s)
	n := len(s)
	if n == 0 {
		return ""
	}
	var b strings.Builder
	if n%k == 0 {
		for i := 0; i < n; i += k {
			b.WriteString(s[i : i+k])
			b.WriteByte('-')
		}
		ans := b.String()
		return ans[:len(ans)-1]
	}
	rem := n % k
	b.WriteString(s[:rem])
	b.WriteByte('-')
	for i := rem; i < n; i += k {
		b.WriteString(s[i : i+k])
		b.WriteByte('-')
	}
	ans := b.String()
	return ans[:len(ans)-1]
}
