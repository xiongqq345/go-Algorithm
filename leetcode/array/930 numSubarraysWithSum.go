package array

// 给你一个二元数组 nums ，和一个整数 goal ，请你统计并返回有多少个和为 goal 的 非空 子数组。
//
//子数组 是数组的一段连续部分。
func numSubarraysWithSum(nums []int, goal int) int {
	mp := make([]int, len(nums)+1)
	var sum, ans int
	for _, v := range nums {
		mp[sum]++
		sum += v
		if sum >= goal {
			ans += mp[sum-goal]
		}
	}
	return ans
}
