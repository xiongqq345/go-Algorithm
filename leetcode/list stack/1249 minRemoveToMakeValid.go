package list_stack

import "strings"

//给你一个由 '('、')' 和小写字母组成的字符串 s。
//
//你需要从字符串中删除最少数目的 '(' 或者 ')' （可以删除任意位置的括号)，使得剩下的「括号字符串」有效。
//
//请返回任意一个合法字符串。
//
//有效「括号字符串」应当符合以下 任意一条 要求：
//
//空字符串或只包含小写字母的字符串
//可以被写作 AB（A 连接 B）的字符串，其中 A 和 B 都是有效「括号字符串」
//可以被写作 (A) 的字符串，其中 A 是一个有效的「括号字符串」
//
func minRemoveToMakeValid(s string) string {
	str := []byte(s)
	var sum int
	for i, ch := range str {
		if ch == '(' {
			sum++
		} else if ch == ')' {
			if sum > 0 {
				sum--
			} else {
				str[i] = '-'
			}
		}
	}
	sum = 0
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == ')' {
			sum++
		} else if str[i] == '(' {
			if sum > 0 {
				sum--
			} else {
				str[i] = '-'
			}
		}
	}
	return strings.ReplaceAll(string(str), "-", "")
}