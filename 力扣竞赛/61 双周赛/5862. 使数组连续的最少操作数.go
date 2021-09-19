package _1_双周赛

import "sort"

func minOperations(nums []int) int {
	var ans int
	n := len(nums)
	sort.Ints(nums)
	nums = unique(nums)
	for r, v := range nums {
		l := sort.SearchInts(nums[:r], v-n+1)
		ans = max(ans, r-l+1) // [l,r] 内的元素均可以保留
	}
	return n - ans
}

func unique(a []int) []int {
	k := 0
	for _, v := range a[1:] {
		if a[k] != v {
			k++
			a[k] = v
		}
	}
	return a[:k+1]
}
