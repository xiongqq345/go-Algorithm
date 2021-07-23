package array

import "sort"

// 给你一个 无重叠的 ，按照区间起始端点排序的区间列表。
//
//在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
//
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/insert-interval
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
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
	var result [][]int
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
