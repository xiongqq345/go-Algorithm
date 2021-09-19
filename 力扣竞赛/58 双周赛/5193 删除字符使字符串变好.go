package _8_双周赛

import "strings"

// 一个字符串如果没有 三个连续 相同字符，那么它就是一个 好字符串 。
//
//给你一个字符串 s ，请你从 s 删除 最少 的字符，使它变成一个 好字符串 。
//
//请你返回删除后的字符串。题目数据保证答案总是 唯一的 。
//
func makeFancyString(s string) string {
	str := []byte(s)
	var candidate byte
	var cnt int
	for i := range str {
		if str[i] != candidate {
			candidate = str[i]
			cnt = 0
		}
		cnt++
		if cnt >= 3 {
			str[i] = '-'
		}
	}
	return strings.ReplaceAll(string(str), "-", "")
}
