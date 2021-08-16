package greedy

import (
	"math"
	"sort"
)

// 给出 n 个数对。 在每一个数对中，第一个数字总是比第二个数字小。
//
//现在，我们定义一种跟随关系，当且仅当 b < c 时，数对(c, d) 才可以跟在 (a, b) 后面。我们用这种形式来构造一个数对链。
//
//给定一个数对集合，找出能够形成的最长数对链的长度。你不需要用到所有的数对，你可以以任何顺序选择其中的一些数对来构造。
//
func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	pre := math.MinInt32
	var cnt int
	for _, pair := range pairs {
		if pair[0] > pre {
			cnt++
			pre = pair[1]
		}
		if pair[1] < pre {
			pre = pair[1]
		}
	}
	return cnt
}
