package design

import "math/rand"

// 给你一个整数数组 nums ，设计算法来打乱一个没有重复元素的数组。
//
//实现 Solution class:
//
//Solution(int[] nums) 使用整数数组 nums 初始化对象
//int[] reset() 重设数组到它的初始状态并返回
//int[] shuffle() 返回数组随机打乱后的结果
//
//

type Solution struct {
	origin []int
}

func Constructor(nums []int) Solution {
	return Solution{
		origin: nums,
	}
}

/** Resets the array to its original configuration and return it. */
func (s *Solution) Reset() []int {
	return s.origin
}

/** Returns a random shuffling of the array. */
func (s *Solution) Shuffle() []int {
	t := make([]int, len(s.origin))
	copy(t, s.origin)
	rand.Shuffle(len(t), func(i, j int) {
		t[i], t[j] = t[j], t[i]
	})
	return t
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
