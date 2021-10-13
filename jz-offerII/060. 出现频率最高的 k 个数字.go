package jz_offerII

import (
	"container/heap"
)

//给定一个整数数组 nums 和一个整数 k ，请返回其中出现频率前 k 高的元素。可以按 任意顺序 返回答案。

type heapInt3 [][2]int

func (h *heapInt3) Less(i, j int) bool {
	return (*h)[i][1] < (*h)[j][1]
}

func (h *heapInt3) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *heapInt3) Push(v interface{}) {
	*h = append(*h, v.([2]int))
}

func (h *heapInt3) Pop() interface{} {
	t := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return t
}

func (h *heapInt3) Len() int {
	return len(*h)
}

func topKFrequent(nums []int, k int) []int {
	occur := make(map[int]int)
	for _, v := range nums {
		occur[v]++
	}
	h := new(heapInt3)
	for num, v := range occur {
		heap.Push(h, [2]int{num, v})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	var ans []int
	for _, v := range *h {
		ans = append(ans, v[0])
	}
	return ans
}
