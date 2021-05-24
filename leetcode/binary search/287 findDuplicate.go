package binary_search

//给定一个包含 n + 1 个整数的数组 nums ，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。
//
//假设 nums 只有 一个重复的整数 ，找出 这个重复的数 。
func findDuplicate(nums []int) int {
	m := make([]int, len(nums))
	for _, v := range nums {
		if m[v] > 0 {
			return v
		}
		m[v]++
	}
	return 0
}
