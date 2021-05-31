package list_stack

import "strconv"

//根据 逆波兰表示法，求表达式的值。
//
//有效的算符包括 +、-、*、/ 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
//
//
//
//说明：
//
//整数除法只保留整数部分。
//给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。

func evalRPN(tokens []string) int {
	var stack []int
	for _, token := range tokens {
		if num, err := strconv.Atoi(token); err == nil {
			stack = append(stack, num)
		} else {
			v1, v2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, v1+v2)
			case "-":
				stack = append(stack, v1-v2)
			case "*":
				stack = append(stack, v1*v2)
			case "/":
				stack = append(stack, v1/v2)
			}
		}
	}
	return stack[0]
}
