package design

import (
	"math/rand"
)

type Solution struct {
	origin []int
}

func Constructor(nums []int) Solution {
	return Solution{origin: nums}
}

func (s *Solution) Reset() []int {
	return s.origin
}

func (s *Solution) Shuffle() []int {
	res := make([]int, len(s.origin))
	copy(res, s.origin)
	rand.Shuffle(len(res), func(i, j int) {
		res[i], res[j] = res[j], res[i]
	})
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
