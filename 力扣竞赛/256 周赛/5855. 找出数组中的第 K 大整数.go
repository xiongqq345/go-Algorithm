package _56_周赛

import (
	"sort"
)

// 给你一个字符串数组 nums 和一个整数 k 。nums 中的每个字符串都表示一个不含前导零的整数。
//
//返回 nums 中表示第 k 大整数的字符串。
//
//注意：重复的数字在统计时会视为不同元素考虑。例如，如果 nums 是 ["1","2","2"]，那么 "2" 是最大的整数，"2" 是第二大的整数，"1" 是第三大的整数。
//

func kthLargestNumber(a []string, k int) string {
	sort.Slice(a, func(i, j int) bool {
		s, t := a[i], a[j]
		if len(s) != len(t) {
			return len(s) > len(t)
		}
		return s > t
	})
	return a[k-1]
}
