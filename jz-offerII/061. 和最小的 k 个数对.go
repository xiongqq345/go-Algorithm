package jz_offerII

import (
	"container/heap"
	"math"
	"sort"
)

//给定两个以升序排列的整数数组 nums1 和 nums2 , 以及一个整数 k 。
//
//定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。
//
//请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。
//
type cusSort [][]int

func (c *cusSort) Push(x interface{}) {
	*c = append(*c, x.([]int))
}

func (c *cusSort) Pop() interface{} {
	length := len(*c)
	val := (*c)[length-1]
	*c = (*c)[:length-1]

	return val
}

func (c *cusSort) Peek() []int {
	return (*c)[0]
}

func (c cusSort) Len() int {
	return len(c)
}

func (c cusSort) Less(i, j int) bool {
	return c[i][0]+c[i][1] > c[j][0]+c[j][1]
}

func (c cusSort) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	h := &cusSort{{nums1[0], nums2[0]}}
	heap.Init(h)
	nums1Min := int(math.Min(float64(k), float64(len(nums1))))
	nums2Min := int(math.Min(float64(k), float64(len(nums2))))
	for i := 0; i < nums1Min; i++ {
		for j := 0; j < nums2Min; j++ {
			if i == 0 && j == 0 {
				continue
			}

			if h.Len() != k || (nums1[i]+nums2[j]) <=
				h.Peek()[0]+h.Peek()[1] {
				if h.Len() == k {
					heap.Pop(h)
				}
				heap.Push(h, []int{nums1[i], nums2[j]})
			}
		}
	}
	sort.Sort(*h)
	sort.Sort(sort.Reverse(h))
	if h.Len() > k {
		return (*h)[:k]
	}
	return *h
}
