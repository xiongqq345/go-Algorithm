package jz_offer

// 用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )
type CQueue struct {
	stackIn, stackOut []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (q *CQueue) AppendTail(value int) {
	q.stackIn = append(q.stackIn, value)
}

func (q *CQueue) DeleteHead() int {
	n := len(q.stackOut)
	if n > 0 {
		popV := q.stackOut[n-1]
		q.stackOut = q.stackOut[:n-1]
		return popV
	}

	for len(q.stackIn) > 0 {
		tmp := q.stackIn[len(q.stackIn)-1]
		q.stackIn = q.stackIn[:len(q.stackIn)-1]
		q.stackOut = append(q.stackOut, tmp)
	}

	n = len(q.stackOut)
	if n > 0 {
		popV := q.stackOut[n-1]
		q.stackOut = q.stackOut[:n-1]
		return popV
	}
	return -1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
