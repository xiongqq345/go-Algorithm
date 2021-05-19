package array

// 给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。
func subarraySum(nums []int, k int) int {
	var count int
	for i := range nums {
		sum := 0
		for j := i; j >= 0; j-- {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}
	return count
}

func subarraySum2(nums []int, k int) int {
	var sum, count int
	m := map[int]int{0: 1}
	for i := range nums {
		sum += nums[i]
		count += m[sum-k]
		m[sum]++
	}
	return count
}
