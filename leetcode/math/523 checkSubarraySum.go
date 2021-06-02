package math

//给定一个包含 非负数 的数组和一个目标 整数 k ，编写一个函数来判断该数组是否含有连续的子数组，其大小至少为 2，且总和为 k 的倍数，即总和为 n * k ，其中 n 也是一个整数。
func checkSubarraySum(nums []int, k int) bool {
	var sum int
	m := make(map[int]int)
	m[0] = -1
	for i, num := range nums {
		sum += num
		if v, ok := m[sum%k]; ok {
			if i-v >= 2 {
				return true
			}
		} else {
			m[sum%k] = i
		}
	}
	return false
}
