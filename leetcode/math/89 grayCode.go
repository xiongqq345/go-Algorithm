package math

// 格雷编码是一个二进制数字系统，在该系统中，两个连续的数值仅有一个位数的差异。
//
//给定一个代表编码总位数的非负整数 n，打印其格雷编码序列。即使有多个不同答案，你也只需要返回其中一种。
//
//格雷编码序列必须以 0 开头。
func grayCode(n int) []int {
	ans := make([]int, 1<<n)
	for i := 1; i <= n; i++ {
		s := 1 << (i - 1)
		for j := 0; j < 1<<(i-1); j++ {
			ans[s+j] = ans[s-1-j] + s
		}
	}
	return ans
}
