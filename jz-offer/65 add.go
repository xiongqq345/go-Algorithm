package jz_offer

// 写一个函数，求两个整数之和，要求在函数体内不得使用 “+”、“-”、“*”、“/” 四则运算符号。
func add(a int, b int) int {
	var carry int
	for b != 0 {
		carry = (a & b) << 1
		a ^= b
		b = carry
	}
	return a
}

func add2(a int, b int) int {
	if b == 0 {
		return a
	}
	return add2(a^b, (a&b)<<1)
}
