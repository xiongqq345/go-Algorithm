package array

import "sort"

//给你一个区间列表，请你删除列表中被其他区间所覆盖的区间。
//
//只有当 c <= a 且 b <= d 时，我们才认为区间 [a,b) 被区间 [c,d) 覆盖。
//
//在完成所有删除操作后，请你返回列表中剩余区间的数目。
func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] > intervals[j][1]
	})

	var remove int
	last := []int{0, 0}
	for _, v := range intervals {
		if v[1] <= last[1] {
			remove++
		} else {
			last = v
		}
	}
	return len(intervals) - remove
}
