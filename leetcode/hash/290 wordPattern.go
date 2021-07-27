package hash

import "strings"

// 给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。
//
//这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。

func wordPattern(pattern string, s string) bool {
	words := strings.Fields(s)
	if len(pattern) != len(words) {
		return false
	}

	mp := make(map[rune]string)
	strMp := make(map[string]bool)
	for i, v := range pattern {
		if _, ok := mp[v]; !ok {
			if strMp[words[i]] {
				return false
			}
			mp[v] = words[i]
			strMp[words[i]] = true
		} else {
			if mp[v] != words[i] {
				return false
			}
		}
	}
	return true
}
