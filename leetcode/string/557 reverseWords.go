package string

import (
	"sort"
	"strings"
)

// 给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
//
//示例：
//
//输入："Let's take LeetCode contest"
//输出："s'teL ekat edoCteeL tsetnoc"
func reverseWords(s string) string {
	words := strings.Fields(s)
	for i := range words {
		words[i] = reverse(words[i])
	}
	return strings.Join(words, " ")
}

func reverse(s string) string {
	str := []byte(s)
	n := len(s) - 1
	for i := range str[:(n+1)/2] {
		str[i], str[n-i] = str[n-i], str[i]
	}
	return string(str)
}
