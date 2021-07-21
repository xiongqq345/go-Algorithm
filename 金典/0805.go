package coding

// 递归乘法。 写一个递归函数，不使用 * 运算符， 实现两个正整数的相乘。可以使用加号、减号、位移，但要吝啬一些。
func multiply(A int, B int) int {
	if B == 1 {
		return A
	}
	return A + multiply(A, B-1)
}
