package jz_offerII

import "strconv"

// 给定两个 01 字符串 a 和 b ，请计算它们的和，并以二进制字符串的形式输出。
//
//输入为 非空 字符串且只包含数字 1 和 0。
func addBinary(a string, b string) string {
	var ans string
	s1, s2 := []byte(a), []byte(b)
	var carry int
	for len(s1) > 0 || len(s2) > 0 {
		var v1, v2 int
		if len(s1) > 0 {
			v1 = int(s1[len(s1)-1]) - 48
			s1 = s1[:len(s1)-1]
		}
		if len(s2) > 0 {
			v2 = int(s2[len(s2)-1]) - 48
			s2 = s2[:len(s2)-1]
		}
		sum := v1 + v2 + carry
		carry, sum = sum/2, sum%2
		ans = strconv.Itoa(sum) + ans
	}
	if carry != 0 {
		ans = strconv.Itoa(carry) + ans
	}
	return ans
}
