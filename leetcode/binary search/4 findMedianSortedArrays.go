package binary_search

// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	k := (m + n) / 2
	i, j := 0, 0
	var pre, cur int
	for i+j <= k {
		if i == m {
			pre, cur = cur, nums2[j]
			j++
			continue
		}
		if j == n {
			pre, cur = cur, nums1[i]
			i++
			continue
		}
		if nums1[i] < nums2[j] {
			pre, cur = cur, nums1[i]
			i++
		} else {
			pre, cur = cur, nums2[j]
			j++
		}
	}

	if (m+n)%2 == 0 {
		return (float64(pre + cur)) / 2
	} else {
		return float64(cur)
	}
}
