package dp

import "sort"

// 给出 n 个数对。 在每一个数对中，第一个数字总是比第二个数字小。
//
//现在，我们定义一种跟随关系，当且仅当 b < c 时，数对(c, d) 才可以跟在 (a, b) 后面。我们用这种形式来构造一个数对链。
//
//给定一个数对集合，找出能够形成的最长数对链的长度。你不需要用到所有的数对，你可以以任何顺序选择其中的一些数对来构造。
//
func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][0] < pairs[j][0] {
			return true
		} else if pairs[i][0] > pairs[j][0] {
			return false
		}
		return pairs[i][1] < pairs[j][1]
	})
	last := pairs[0][1]
	cnt := 1
	for i := 1; i < len(pairs); i++ {
		if pairs[i][0] > last {
			cnt++
			last = pairs[i][1]
		}
	}
	return cnt
}
