package dp

// 数组 A 包含 N 个数，且索引从0开始。数组 A 的一个子数组划分为数组 (P, Q)，P 与 Q 是整数且满足 0<=P<Q<N 。
//
//如果满足以下条件，则称子数组(P, Q)为等差数组：
//
//元素 A[P], A[p + 1], ..., A[Q - 1], A[Q] 是等差的。并且 P + 1 < Q 。
//
//函数要返回数组 A 中所有为等差数组的子数组个数。
func numberOfArithmeticSlices(nums []int) int {
	var dp, sum int
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp++
			sum += dp
		} else {
			dp = 0
		}
	}
	return sum
}
