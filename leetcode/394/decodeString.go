package _94

import (
	"container/list"
	"strings"
)

func decodeString(s string) string {
	list.New()
	res, multi := "", 0
	mulStack := make([]int, 0)
	strStack := make([]string, 0)
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] >= '0' && s[i] <= '9':
			multi = multi*10 + int(s[i]) - '0'
		case s[i] == '[':
			mulStack = append(mulStack, multi)
			strStack = append(strStack, res)
			multi, res = 0, ""
		case s[i] == ']':
			res = strStack[len(strStack)-1] + strings.Repeat(res, mulStack[len(mulStack)-1])
			strStack, mulStack = strStack[:len(strStack)-1], mulStack[:len(mulStack)-1]
		default:
			res += string(s[i])
		}
	}
	return res
}
