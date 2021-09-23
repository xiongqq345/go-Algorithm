package jz_offerII

import (
	"sort"
	"strconv"
)

//给定一个 24 小时制（小时:分钟 "HH:MM"）的时间列表，找出列表中任意两个时间的最小时间差并以分钟数表示。
func findMinDifference(timePoints []string) int {
	arr := make([]int, len(timePoints))
	for i, v := range timePoints {
		arr[i] = atoi(v[:2])*60 + atoi(v[3:5])
	}
	sort.Ints(arr)
	ans := arr[0] + 60*24 - arr[len(arr)-1]
	for i := range arr[1:] {
		ans = min(ans, arr[i+1]-arr[i])
	}
	return ans
}

func atoi(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
