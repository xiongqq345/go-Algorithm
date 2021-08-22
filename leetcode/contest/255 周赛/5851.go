package _55_周赛

import (
	"sort"
	"strconv"
)

// 给你一个字符串数组 nums ，该数组由 n 个 互不相同 的二进制字符串组成，且每个字符串长度都是 n 。请你找出并返回一个长度为 n 且 没有出现 在 nums 中的二进制字符串。如果存在多种答案，只需返回 任意一个 即可。
func findDifferentBinaryString(arr []string) string {
	mp := make(map[string]bool)
	var nums []string
	for _, v := range arr {
		if !mp[v] {
			nums = append(nums, v)
		}
		mp[v] = true
	}
	sort.Strings(nums)
	idx := sort.Search(len(nums), func(i int) bool {
		base, _ := strconv.ParseInt(nums[i], 2, 16)
		return base != int64(i)
	})
	ans := strconv.FormatInt(int64(idx), 2)
	diff := len(arr) - len(ans)
	for diff > 0 {
		ans = "0" + ans
		diff--
	}
	return ans
}
