package leetcode

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
