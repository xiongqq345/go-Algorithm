package dp

import "sort"

// 给出一个单词列表，其中每个单词都由小写英文字母组成。
//
//如果我们可以在 word1 的任何地方添加一个字母使其变成 word2，那么我们认为 word1 是 word2 的前身。例如，"abc" 是 "abac" 的前身。
//
//词链是单词 [word_1, word_2, ..., word_k] 组成的序列，k >= 1，其中 word_1 是 word_2 的前身，word_2 是 word_3 的前身，依此类推。
//
//从给定单词列表 words 中选择单词组成词链，返回词链的最长可能长度。
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-string-chain
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	mp := make(map[string]int)
	var ans int
	for _, w := range words {
		for i := range w {
			t := w[:i] + w[i+1:]
			mp[w] = max(mp[w], mp[t]+1)
			ans = max(ans, mp[w])
		}
	}
	return ans
}
