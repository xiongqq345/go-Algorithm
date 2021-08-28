package jz_offerII

// 给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
func findMaxLength(nums []int) int {
	mp := make(map[int]int)
	mp[0] = -1
	var sum, ans int
	for i := range nums {
		if nums[i] == 0 {
			sum--
		} else {
			sum++
		}
		if v, ok := mp[sum]; !ok {
			mp[sum] = i
		} else {
			ans = max(ans, i-v)
		}
	}
	return ans
}
