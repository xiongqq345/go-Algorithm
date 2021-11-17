package bit

//给定一个字符串数组 words，找到 length(word[i]) * length(word[j]) 的最大值，并且这两个单词不含有公共字母。你可以认为每个单词只包含小写字母。如果不存在这样的两个单词，返回 0。
//
func maxProduct(words []string) int {
	masks := make([]int, len(words))
	for i, word := range words {
		for _, ch := range word {
			masks[i] |= 1 << (ch - 'a')
		}
	}

	var ans int
	for i, x := range masks {
		for j, y := range masks[:i] {
			if len(words[i])*len(words[j]) > ans && x&y == 0 {
				ans = len(words[i]) * len(words[j])
			}
		}
	}
	return ans
}
