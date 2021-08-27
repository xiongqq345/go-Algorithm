package jz_offerII

// 给定一个整数数组和一个整数 k ，请找到该数组中和为 k 的连续子数组的个数。
func subarraySum(nums []int, k int) int {
	var sum, cnt int
	mp := map[int]int{0: 1}
	for _, num := range nums {
		sum += num
		cnt += mp[sum-k]
		mp[sum]++
	}
	return cnt
}
