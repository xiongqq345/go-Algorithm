package priority_queue

type heapInt []int

func (h *heapInt) Len() int { return len(*h) }

func (h *heapInt) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *heapInt) Less(i, j int) bool { return (*h)[i] < (*h)[j] }

func (h *heapInt) Pop() interface{} {
	t := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return t
}

func (h *heapInt) Push(v interface{}) {
	*h = append(*h, v.(int))
}
