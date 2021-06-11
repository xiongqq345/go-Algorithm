package jz_offer

import "strconv"

// 给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。
//
func translateNum(num int) int {
	str := strconv.Itoa(num)
	dp := make([]int, len(str)+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= len(str); i++ {
		substr := str[i-2 : i]
		if substr <= "25" && substr >= "10" {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(str)]
}

func translateNum2(num int) int {
	str := strconv.Itoa(num)
	p, q := 1, 1
	for i := 2; i <= len(str); i++ {
		substr := str[i-2 : i]
		if substr <= "25" && substr >= "10" {
			p, q = q, p+q
		} else {
			p = q
		}
	}
	return q
}
