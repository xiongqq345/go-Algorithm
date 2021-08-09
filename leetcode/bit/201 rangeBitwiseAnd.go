package bit

// 给你两个整数 left 和 right ，表示区间 [left, right] ，返回此区间内所有数字 按位与 的结果（包含 left 、right 端点）。
func rangeBitwiseAnd(m int, n int) int {
	for m < n {
		n &= n - 1
	}
	return n
}
