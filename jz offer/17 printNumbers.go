package jz_offer

// 输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。
func printNumbers(n int) []int {
	length := 1
	for n > 0 {
		length *= 10
		n--
	}
	ans := make([]int, length-1)
	for i := range ans {
		ans[i] = i + 1
	}
	return ans
}
