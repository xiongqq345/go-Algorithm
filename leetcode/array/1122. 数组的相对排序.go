package array

import "sort"

//给你两个数组，arr1 和 arr2，
//
//arr2 中的元素各不相同
//arr2 中的每个元素都出现在 arr1 中
//对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。未在 arr2 中出现过的元素需要按照升序放在 arr1 的末尾。
//
func relativeSortArray(arr1 []int, arr2 []int) []int {
	mp := make(map[int]int)
	for i, v := range arr2 {
		mp[v] = i
	}
	sort.Slice(arr1, func(i, j int) bool {
		x, y := arr1[i], arr1[j]
		ix, okx := mp[x]
		iy, oky := mp[y]
		if okx && oky {
			return ix < iy
		}
		if okx || oky {
			return okx
		}
		return x < y
	})
	return arr1
}
