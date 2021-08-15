package jz_offerII

//给定一个非负整数 n ，请计算 0 到 n 之间的每个数字的二进制表示中 1 的个数，并输出一个数组。
func countBits(num int) []int {
	ans := make([]int, num+1)
	for i := range ans {
		ans[i] = ans[i>>1] + i&1
	}
	return ans
}
