package _53_周赛

import "container/heap"

// 给你一个整数数组 piles ，数组 下标从 0 开始 ，其中 piles[i] 表示第 i 堆石子中的石子数量。另给你一个整数 k ，请你执行下述操作 恰好 k 次：
//
//选出任一石子堆 piles[i] ，并从中 移除 floor(piles[i] / 2) 颗石子。
//注意：你可以对 同一堆 石子多次执行此操作。
//
//返回执行 k 次操作后，剩下石子的 最小 总数。
//
//floor(x) 为 小于 或 等于 x 的 最大 整数。（即，对 x 向下取整）。
//
type heapInt []int

//Less  小于就是小跟堆，大于号就是大根堆
func (h *heapInt) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *heapInt) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *heapInt) Len() int           { return len(*h) }
func (h *heapInt) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *heapInt) Pop() interface{} {
	t := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return t
}
func (h *heapInt) Peek() int {
	return (*h)[0]
}

func minStoneSum(piles []int, k int) int {
	h := new(heapInt)
	for _, v := range piles {
		heap.Push(h, v)
	}
	for k > 0 {
		pile := heap.Pop(h).(int)
		heap.Push(h, pile-pile/2)
		k--
	}
	var sum int
	for _, p := range *h {
		sum += p
	}
	return sum
}
