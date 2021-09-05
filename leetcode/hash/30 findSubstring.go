package hash

// 给定一个字符串 s 和一些 长度相同 的单词 words 。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
//
//注意子串要与 words 中的单词完全匹配，中间不能有其他字符 ，但不需要考虑 words 中单词串联的顺序。

func findSubstring(s string, words []string) []int {
	var ans []int
	step := len(words[0])
	var length int
	mp := make(map[string]int)
	for _, w := range words {
		mp[w]++
		length += len(w)
	}

p:
	for i := 0; i <= len(s)-length; i++ {
		tmp := make(map[string]int)
		last := i
		for idx := i; idx <= i+length; idx += step {
			if _, ok := mp[s[last:idx]]; ok {
				tmp[s[last:idx]]++
				if tmp[s[last:idx]] > mp[s[last:idx]] {
					continue p
				}
				last = idx
			}
		}
		if compareMap(mp, tmp) {
			ans = append(ans, i)
		}
	}
	return ans
}

func compareMap(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, w := range m1 {
		if m2[k] != w {
			return false
		}
	}
	return true
}
