package jz_offerII

import "math"

// 给定两个整数 a 和 b ，求它们的除法的商 a/b ，要求不得使用乘号 '*'、除号 '/' 以及求余符号 '%' 。
func divide(dividend int, divisor int) int {
	a, b := dividend, divisor
	if a == 0 {
		return 0
	}
	if b == 1 {
		return a
	}
	if b == -1 {
		if a > math.MinInt32 {
			return -a
		}
		return math.MaxInt32
	}
	sign := 1
	if a > 0 && b < 0 || a < 0 && b > 0 {
		sign = -1
	}
	a = abs(a)
	b = abs(b)
	ans := div(a, b)
	if sign > 0 {
		return ans
	}
	return -ans
}

func div(a, b int) int {
	if a < b {
		return 0
	}
	cnt, tmp := 1, b
	for tmp+tmp < a {
		tmp *= 2
		cnt *= 2
	}
	return cnt + div(a-tmp, b)
}
