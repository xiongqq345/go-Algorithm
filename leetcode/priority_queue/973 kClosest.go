package priority_queue

import "container/heap"

// 我们有一个由平面上的点组成的列表 points。需要从中找出 K 个距离原点 (0, 0) 最近的点。
//
//（这里，平面上两点之间的距离是欧几里德距离。）
//
//你可以按任何顺序返回答案。除了点坐标的顺序之外，答案确保是唯一的。

type heapInt2 [][2]int

func (h heapInt2) Less(i, j int) bool {
	return h[i][0]*h[i][0]+h[i][1]*h[i][1] >
		h[j][0]*h[j][0]+h[j][1]*h[j][1]
}
func (h *heapInt2) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *heapInt2) Len() int      { return len(*h) }

func (h *heapInt2) Pop() interface{} {
	t := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return t
}
func (h *heapInt2) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func kClosest(points [][]int, k int) [][]int {
	h := new(heapInt2)
	for _, p := range points {
		heap.Push(h, [2]int{p[0], p[1]})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	var ans [][]int
	for i := 0; i < k; i++ {
		v := heap.Pop(h).([2]int)
		ans = append(ans, []int{v[0], v[1]})
	}
	return ans
}
