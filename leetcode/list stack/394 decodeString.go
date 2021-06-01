package list_stack

import "strings"

//给定一个经过编码的字符串，返回它解码后的字符串。
//
//编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
//
//你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
//
//此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像3a或2[4]的输入。
func decodeString(s string) string {
	var multi int
	var res string
	var multiStack []int
	var strStack []string
	for _, c := range s {
		switch {
		case c >= '0' && c <= '9':
			multi = multi*10 + int(c) - '0'
		case c == '[':
			multiStack = append(multiStack, multi)
			strStack = append(strStack, res)
			multi, res = 0, ""
		case c == ']':
			res = strStack[len(strStack)-1] + strings.Repeat(res, multiStack[len(multiStack)-1])
			strStack, multiStack = strStack[:len(strStack)-1], multiStack[:len(multiStack)-1]
		default:
			res += string(c)
		}
	}
	return res
}
