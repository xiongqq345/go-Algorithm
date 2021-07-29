package list_stack

// 给你两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。
//
//请你找出 nums1 中每个元素在 nums2 中的下一个比其大的值。
//
//nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出 -1 。
//
//
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	mp := make(map[int]int, len(nums2))
	var stack []int
	for _, num := range nums2 {
		for len(stack) > 0 && num > stack[len(stack)-1] {
			mp[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}

	ans := make([]int, 0, len(nums1))
	for _, num := range nums1 {
		if v, ok := mp[num]; !ok {
			ans = append(ans, -1)
		} else {
			ans = append(ans, v)
		}
	}
	return ans
}
