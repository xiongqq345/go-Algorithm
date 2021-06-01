package greedy

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	count, pre := 1, intervals[0][1]
	for _, interval := range intervals {
		if interval[0] >= pre {
			count++
			pre = interval[1]
		}
	}
	return n - count
}
