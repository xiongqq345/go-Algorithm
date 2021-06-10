package list_stack

//输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。例如，序列 {1,2,3,4,5} 是某栈的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。
func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	for _, num := range pushed {
		stack = append(stack, num)
		for len(stack) > 0 && popped[0] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			popped = popped[1:]
		}
	}
	return len(popped) == 0
}
