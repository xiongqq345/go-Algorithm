package dp

import "math/bits"

// 给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。
//func countBits(num int) []int {
//	bits := make([]int, num+1)
//	for i := 1; i <= num; i++ {
//		bits[i] = bits[i>>1] + i&1
//	}
//	return bits
//}

func countBits(num int) []int {
	ans := make([]int, num+1)
	for i := range ans {
		ans[i] = bits.OnesCount(uint(i))
	}
	return ans
}
