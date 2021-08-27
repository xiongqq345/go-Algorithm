package jz_offerII

import "strings"

// 给定一个字符串数组 words，请计算当两个字符串 words[i] 和 words[j] 不包含相同字符时，它们长度的乘积的最大值。假设字符串中只包含英语的小写字母。如果没有不包含相同字符的一对字符串，返回 0。
func maxProduct(words []string) int {
	var ans int
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if !strings.ContainsAny(words[i], words[j]) && ans < len(words[i])*len(words[j]) {
				ans = len(words[i]) * len(words[j])
			}
		}
	}
	return ans
}
