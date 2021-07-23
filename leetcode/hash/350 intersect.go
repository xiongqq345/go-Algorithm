package hash

// 给定两个数组，编写一个函数来计算它们的交集。
func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	var ans []int
	for _, v := range nums1 {
		m[v]++
	}
	for _, v := range nums2 {
		if m[v] > 0 {
			m[v]--
			ans = append(ans, v)
		}
	}
	return ans
}
