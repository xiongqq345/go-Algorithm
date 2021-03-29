package leetcode

import "sort"

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	return merge(intervals)
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := make([][]int, 0)
	lens := len(intervals)
	for i := 0; i < lens; {
		if i == lens-1 {
			result = append(result, intervals[i])
			break
		}
		j := i + 1
		for j < lens && intervals[i][1] >= intervals[j][0] {
			if intervals[i][1] < intervals[j][1] {
				intervals[i][1] = intervals[j][1]
			}
			j++
		}

		result = append(result, intervals[i])
		i = j
	}
	return result
}
