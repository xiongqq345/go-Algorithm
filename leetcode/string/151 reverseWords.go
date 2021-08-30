package string

import "strings"

//给你一个字符串 s ，逐个翻转字符串中的所有 单词 。
//
//单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
//
//请你返回一个翻转 s 中单词顺序并用单个空格相连的字符串。
//
//说明：
//
//输入字符串 s 可以在前面、后面或者单词间包含多余的空格。
//翻转后单词间应当仅用一个空格分隔。
//翻转后的字符串中不应包含额外的空格。
//
func reverseWords(s string) string {
	fields := strings.Fields(s)
	n := len(fields)
	for i := 0; i < n/2; i++ {
		fields[i], fields[n-1-i] = fields[n-1-i], fields[i]
	}
	return strings.Join(fields, " ")
}
