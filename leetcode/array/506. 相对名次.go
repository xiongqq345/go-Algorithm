package array

import (
	"sort"
	"strconv"
)

//给你一个长度为 n 的整数数组 score ，其中 score[i] 是第 i 位运动员在比赛中的得分。所有得分都 互不相同 。
//
//运动员将根据得分 决定名次 ，其中名次第 1 的运动员得分最高，名次第 2 的运动员得分第 2 高，依此类推。运动员的名次决定了他们的获奖情况：
//
//名次第 1 的运动员获金牌 "Gold Medal" 。
//名次第 2 的运动员获银牌 "Silver Medal" 。
//名次第 3 的运动员获铜牌 "Bronze Medal" 。
//从名次第 4 到第 n 的运动员，只能获得他们的名次编号（即，名次第 x 的运动员获得编号 "x"）。
//使用长度为 n 的数组 answer 返回获奖，其中 answer[i] 是第 i 位运动员的获奖情况。

var desc = [3]string{"Gold Medal", "Silver Medal", "Bronze Medal"}

func findRelativeRanks(score []int) []string {
	n := len(score)
	type pair struct{ score, idx int }
	arr := make([]pair, n)
	for i, s := range score {
		arr[i] = pair{s, i}
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i].score > arr[j].score })

	ans := make([]string, n)
	for i, p := range arr {
		if i < 3 {
			ans[p.idx] = desc[i]
		} else {
			ans[p.idx] = strconv.Itoa(i + 1)
		}
	}
	return ans
}
