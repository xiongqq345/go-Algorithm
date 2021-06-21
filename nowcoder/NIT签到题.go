//链接：https://ac.nowcoder.com/acm/contest/11191/A
//来源：牛客网
//
//NIT曾经做过一道签到题：输入 aa 和 bb ，输出任意一个 aa 的倍数 cc ，并且保证 bb 是 cc 的倍数 (a<c<b)(a<c<b)
//以NIT现在国家队的实力，做这样的题实在是太侮辱他的智商了，于是他思考着加强这道题目。
//
//输入四个数 a,b,c,da,b,c,d 请你输出一个正整数 xx ，满足 xx 是 aa 的倍数且 xx 是 cc 的倍数，另外满足 bb 是 xx 的倍数且 dd 是 xx 的倍数。 (a<x<b,c<x<d)(a<x<b,c<x<d)

package main

import (
	"fmt"
)

func main() {
	a, b, c, d := 0, 0, 0, 0
	fmt.Scan(&a, &b, &c, &d)
	start := lcm(a, c)
	lcm := start
	if start == a || start == c {
		start *= 2
	}
	end := gcd(b, d)
	if end == b || end == d {
		end--
	}
	var i int
	for i = start; i <= end; i += lcm {
		if b%i == 0 && d%i == 0 {
			fmt.Printf("%d\n", i)
			return
		}
	}
	fmt.Printf("%d\n", -1)
}

func gcd(a, b int) int {
	for a != b {
		if a > b {
			a = a - b
		} else if a < b {
			b = b - a
		}
	}
	return a
}

func lcm(x, y int) int {
	return x * y / gcd(x, y)
}
