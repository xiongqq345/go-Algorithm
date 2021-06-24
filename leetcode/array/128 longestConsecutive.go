package array

// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//
//
//
//进阶：你可以设计并实现时间复杂度为 O(n) 的解决方案吗？
//
//
//
//示例 1：
//
//输入：nums = [100,4,200,1,3,2]
//输出：4
//解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
//示例 2：
//
//输入：nums = [0,3,7,2,5,8,4,6,0,1]
//输出：9
func longestConsecutive(nums []int) int {
	mp := make(map[int]bool)
	for _, v := range nums {
		mp[v] = true
	}
	var maxLen int
	for k := range mp {
		if mp[k-1] {
			continue
		}
		var length int
		for mp[k] {
			length++
			k++
		}
		maxLen = max(maxLen, length)
	}
	return maxLen
}
