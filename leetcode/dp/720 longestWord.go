package dp

import (
	"sort"
)

// 给出一个字符串数组words组成的一本英语词典。从中找出最长的一个单词，该单词是由words词典中其他单词逐步添加一个字母组成。若其中有多个可行的答案，则返回答案中字典序最小的单词。
//
//若无答案，则返回空字符串。
func longestWord(words []string) string {
	sort.Strings(words)
	mp := make(map[string]bool)
	var ans string
	for _, w := range words {
		if len(w) == 1 || mp[w[:len(w)-1]] {
			mp[w] = true
			if len(w) > len(ans) {
				ans = w
			}
		}
	}
	return ans
}
