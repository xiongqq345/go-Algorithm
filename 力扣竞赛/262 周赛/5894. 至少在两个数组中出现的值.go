package _62_周赛

//给你三个整数数组 nums1、nums2 和 nums3 ，请你构造并返回一个 不同 数组，且由 至少 在 两个 数组中出现的所有值组成。数组中的元素可以按 任意 顺序排列。
func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	ansset := make(map[int]bool)
	set := make(map[int]bool)
	for _, v := range nums1 {
		set[v] = true
	}
	for _, v := range nums2 {
		if set[v] {
			ansset[v] = true
		}
	}
	for _, v := range nums2 {
		set[v] = true
	}
	for _, v := range nums3 {
		if set[v] {
			ansset[v] = true
		}
	}
	var ans []int
	for k := range ansset {
		ans = append(ans, k)
	}
	return ans
}
