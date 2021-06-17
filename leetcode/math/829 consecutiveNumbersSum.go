package math

// 给定一个正整数 N，试求有多少组连续正整数满足所有数字之和为 N?
func consecutiveNumbersSum(N int) int {
	res, i := 0, 1
	for N > 0 {
		if N%i == 0 {
			res++
		}
		N -= i
		i++
	}
	return res
}
