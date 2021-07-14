package coding

// 三合一。描述如何只用一个数组来实现三个栈。
//
//你应该实现push(stackNum, value)、pop(stackNum)、isEmpty(stackNum)、peek(stackNum)方法。stackNum表示栈下标，value表示压入的值。
//
//构造函数会传入一个stackSize参数，代表每个栈的大小。
//

type TripleInOne struct {
	stack []int
	cnt   [3]int
	size  int
}

func Constructor(stackSize int) TripleInOne {
	return TripleInOne{
		stack: make([]int, stackSize*3),
		size:  stackSize,
	}
}

func (s *TripleInOne) Push(stackNum int, value int) {
	if s.cnt[stackNum] < s.size {
		s.stack[stackNum*s.size+s.cnt[stackNum]] = value
		s.cnt[stackNum]++
	}
}

func (s *TripleInOne) Pop(stackNum int) int {
	if s.IsEmpty(stackNum) {
		return -1
	}
	s.cnt[stackNum]--
	return s.stack[stackNum*s.size+s.cnt[stackNum]]
}

func (s *TripleInOne) Peek(stackNum int) int {
	if s.IsEmpty(stackNum) {
		return -1
	}
	return s.stack[stackNum*s.size+s.cnt[stackNum]-1]
}

func (s *TripleInOne) IsEmpty(stackNum int) bool {
	return s.cnt[stackNum] == 0
}

/**
 * Your TripleInOne object will be instantiated and called as such:
 * obj := Constructor(stackSize);
 * obj.Push(stackNum,value);
 * param_2 := obj.Pop(stackNum);
 * param_3 := obj.Peek(stackNum);
 * param_4 := obj.IsEmpty(stackNum);
 */
