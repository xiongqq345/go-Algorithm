package array

import "sort"

func merge1(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var ans [][]int
	merged := intervals[0]
	for _, interval := range intervals {
		if interval[0] > merged[1] {
			ans = append(ans, merged)
			merged = interval
		} else if interval[1] > merged[1] {
			merged[1] = interval[1]
		}
	}
	return append(ans, merged)
}
