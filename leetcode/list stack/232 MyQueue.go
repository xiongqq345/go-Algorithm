package list_stack

type MyQueue struct {
	queue []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (q *MyQueue) Push(x int) {
	q.queue = append(q.queue, x)
}

/** Removes the element from in front of queue and returns that element. */
func (q *MyQueue) Pop() int {
	v := q.queue[0]
	q.queue = q.queue[1:]
	return v
}

/** Get the front element. */
func (q *MyQueue) Peek() int {
	return q.queue[0]
}

/** Returns whether the queue is empty. */
func (q *MyQueue) Empty() bool {
	return len(q.queue) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
