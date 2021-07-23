package hash

//给定两个数组，编写一个函数来计算它们的交集。
func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]struct{})
	var ans []int
	for _, v := range nums1 {
		m[v] = struct{}{}
	}
	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			delete(m, v)
			ans = append(ans, v)
		}
	}
	return ans
}
