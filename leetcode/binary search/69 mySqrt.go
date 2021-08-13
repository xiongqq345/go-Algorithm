package binary_search

import "sort"

// 实现int sqrt(int x)函数。
//
//计算并返回x的平方根，其中x 是非负整数。
//
//由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
func mySqrt(x int) int {
	l, r, ans := 0, x, 0
	for l <= r {
		mid := (l + r) / 2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

func mySqrt2(x int) int {
	i := sort.Search(x/2+1, func(i int) bool {
		return i*i > x
	})
	if i*i == x {
		return i
	}
	return i - 1
}
