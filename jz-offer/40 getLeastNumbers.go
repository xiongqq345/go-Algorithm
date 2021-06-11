package jz_offer

import "container/heap"

// 输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

type heapInt []int

//Less  小于就是小跟堆，大于号就是大根堆
func (h *heapInt) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
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

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return nil
	}

	h := new(heapInt)
	for _, num := range arr {
		if h.Len() < k {
			heap.Push(h, num)
		} else {
			if num < h.Peek() {
				heap.Pop(h)
				heap.Push(h, num)
			}
		}
	}
	return *h
}
