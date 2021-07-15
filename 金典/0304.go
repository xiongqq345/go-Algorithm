package coding

type MyQueue struct {
	stackIn, stackOut []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (q *MyQueue) Push(x int) {
	q.stackIn = append(q.stackIn, x)
}

func (q *MyQueue) PushOut(x int) {
	q.stackOut = append(q.stackOut, x)
}

// 倾倒
func (q *MyQueue) in2out() {
	for len(q.stackIn) > 0 {
		q.PushOut(q.stackIn[len(q.stackIn)-1])
		q.stackIn = q.stackIn[:len(q.stackIn)-1]
	}
}

/** Removes the element from in front of queue and returns that element. */
func (q *MyQueue) Pop() int {
	if len(q.stackOut) == 0 {
		q.in2out()
	}
	val := q.stackOut[len(q.stackOut)-1]
	q.stackOut = q.stackOut[:len(q.stackOut)-1]
	return val
}

/** Get the front element. */
func (q *MyQueue) Peek() int {
	if len(q.stackOut) == 0 {
		q.in2out()
	}
	return q.stackOut[len(q.stackOut)-1]
}

/** Returns whether the queue is empty. */
func (q *MyQueue) Empty() bool {
	return len(q.stackIn)+len(q.stackOut) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
