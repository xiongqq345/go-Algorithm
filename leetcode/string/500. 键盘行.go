package string

import "strings"

//给你一个字符串数组 words ，只返回可以使用在 美式键盘 同一行的字母打印出来的单词。键盘如下图所示。
func findWords(words []string) []string {
	var ans []string
	arr := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
P:
	for _, word := range words {
		lower := strings.ToLower(word)
		var idx int
		for i, row := range arr {
			if strings.Index(row, string(lower[0])) >= 0 {
				idx = i
				break
			}
		}

		for _, ch := range lower {
			if strings.Index(arr[idx], string(ch)) == -1 {
				continue P
			}
		}
		ans = append(ans, word)
	}
	return ans
}
