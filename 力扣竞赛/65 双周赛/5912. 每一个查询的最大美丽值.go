package _5_双周赛

import "sort"

//给你一个二维整数数组 items ，其中 items[i] = [pricei, beautyi] 分别表示每一个物品的 价格 和 美丽值 。
//
//同时给你一个下标从 0 开始的整数数组 queries 。对于每个查询 queries[j] ，你想求出价格小于等于 queries[j] 的物品中，最大的美丽值 是多少。如果不存在符合条件的物品，那么查询的结果为 0 。
//
//请你返回一个长度与 queries 相同的数组 answer，其中 answer[j]是第 j 个查询的答案。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/most-beautiful-item-for-each-query
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func maximumBeauty(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool {
		if items[i][0] != items[j][0] {
			return items[i][0] < items[j][0]
		}
		return items[i][1] < items[j][1]
	})
	mp := make(map[int]int)
	var last int
	var arr []int
	for _, v := range items {
		mp[v[0]] = max(mp[last], v[1])
		if v[0] != last {
			arr = append(arr, v[0])
			last = v[0]
		}
	}
	ans := make([]int, len(queries))
	for i, v := range queries {
		idx := sort.SearchInts(arr, v)
		if len(arr) == idx || arr[idx] != v {
			if idx == 0 {
				ans[i] = 0
				continue
			}
			idx--
		}
		ans[i] = mp[arr[idx]]
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
