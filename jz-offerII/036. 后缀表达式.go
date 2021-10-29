package jz_offerII

import "strconv"

//根据 逆波兰表示法，求该后缀表达式的计算结果。
//
//有效的算符包括 +、-、*、/ 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
//
//
//
//说明：
//
//整数除法只保留整数部分。
//给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
//
func evalRPN(tokens []string) int {
	var stack []int
	for _, t := range tokens {
		if num, err := strconv.Atoi(t); err == nil {
			stack = append(stack, num)
		} else {
			n := len(stack)
			v1, v2 := stack[n-2], stack[n-1]
			stack = stack[:n-2]
			switch t {
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
