package priority_queue

import (
	"container/heap"
)

// 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
type heapInt2 [][2]int

func (h *heapInt2) Less(i, j int) bool { return (*h)[i][1] < (*h)[j][1] }
func (h *heapInt2) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *heapInt2) Len() int           { return len(*h) }

func (h *heapInt2) Pop() interface{} {
	t := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return t
}
func (h *heapInt2) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func topKFrequent(nums []int, k int) []int {
	occur := make(map[int]int)
	for _, num := range nums {
		occur[num]++
	}
	h := new(heapInt2)
	for key, v := range occur {
		heap.Push(h, [2]int{key, v})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	var ans []int
	for i := 0; i < k; i++ {
		ans = append(ans, heap.Pop(h).([2]int)[0])
	}
	return ans
}
