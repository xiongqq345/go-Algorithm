package design

import "container/heap"

//设计一个找到数据流中第 k 大元素的类（class）。注意是排序后的第 k 大元素，不是第 k 个不同的元素。
//
//请实现 KthLargest 类：
//
//KthLargest(int k, int[] nums) 使用整数 k 和整数流 nums 初始化对象。
//int add(int val) 将 val 插入数据流 nums 后，返回当前数据流中第 k 大的元素。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/kth-largest-element-in-a-stream
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
type heapInt []int

//Less  小于就是小根堆，大于号就是大根堆
func (h *heapInt) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *heapInt) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *heapInt) Len() int           { return len(*h) }
func (h *heapInt) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *heapInt) Pop() interface{}   { t := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return t }
func (h *heapInt) Peek() int          { return (*h)[0] }

type KthLargest struct {
	heap heapInt
	size int
}

func Constructor(k int, nums []int) KthLargest {
	h := heapInt(nums)
	heap.Init(&h)
	for h.Len() > k {
		heap.Pop(&h)
	}
	return KthLargest{
		size: k,
		heap: h,
	}
}

func (s *KthLargest) Add(val int) int {
	heap.Push(&s.heap, val)
	if len(s.heap) > s.size {
		heap.Pop(&s.heap)
	}
	return s.heap.Peek()
}
