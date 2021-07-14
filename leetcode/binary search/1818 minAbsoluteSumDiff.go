package binary_search

import "sort"

// 给你两个正整数数组 nums1 和 nums2 ，数组的长度都是 n 。
//
//数组 nums1 和 nums2 的 绝对差值和 定义为所有 |nums1[i] - nums2[i]|（0 <= i < n）的 总和（下标从 0 开始）。
//
//你可以选用 nums1 中的 任意一个 元素来替换 nums1 中的 至多 一个元素，以 最小化 绝对差值和。
//
//在替换数组 nums1 中最多一个元素 之后 ，返回最小绝对差值和。因为答案可能很大，所以需要对 109 + 7 取余 后返回。

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	tmp := append(sort.IntSlice{}, nums1...)
	tmp.Sort()
	sum, maxn, n := 0, 0, len(nums1)
	for i, v := range nums2 {
		diff := abs(nums1[i] - v)
		sum += diff
		j := tmp.Search(v)
		if j < n {
			maxn = max(maxn, diff-(tmp[j]-v))
		}
		if j > 0 {
			maxn = max(maxn, diff-(v-tmp[j-1]))
		}
	}
	return (sum - maxn) % (1e9 + 7)
}
