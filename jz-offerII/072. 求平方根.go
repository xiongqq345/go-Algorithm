package jz_offerII

import "sort"

//给定一个非负整数 x ，计算并返回 x 的平方根，即实现 int sqrt(int x) 函数。
//
//正数的平方根有两个，只输出其中的正数平方根。
//
//如果平方根不是整数，输出只保留整数的部分，小数部分将被舍去。
func mySqrt(x int) int {
	return sort.Search(x+1, func(i int) bool {
		return i*i > x
	}) - 1
}
