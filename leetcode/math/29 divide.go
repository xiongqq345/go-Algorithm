package math

import "math"

//给定两个整数，被除数dividend和除数divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
//
//返回被除数dividend除以除数divisor得到的商。
//
//整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2
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
