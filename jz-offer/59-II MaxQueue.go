package jz_offer

// 请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。
//
//若队列为空，pop_front 和 max_value 需要返回 -1
type MaxQueue struct {
	queue    []int
	maxQueue []int
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (q *MaxQueue) Max_value() int {
	if len(q.queue) == 0 {
		return -1
	}
	return q.maxQueue[0]
}

func (q *MaxQueue) Push_back(value int) {
	q.queue = append(q.queue, value)
	for len(q.maxQueue) > 0 && q.maxQueue[len(q.maxQueue)-1] < value {
		q.maxQueue = q.maxQueue[:len(q.maxQueue)-1]
	}
	q.maxQueue = append(q.maxQueue, value)
}

func (q *MaxQueue) Pop_front() int {
	if len(q.queue) == 0 {
		return -1
	}
	front := q.queue[0]
	q.queue = q.queue[1:]
	if q.maxQueue[0] == front {
		q.maxQueue = q.maxQueue[1:]
	}
	return front
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
