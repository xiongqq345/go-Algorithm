package string

import (
	"strconv"
	"strings"
)

//给你一个整数 n ，找出从 1 到 n 各个整数的 Fizz Buzz 表示，并用字符串数组 answer（下标从 1 开始）返回结果，其中：
//
//answer[i] == "FizzBuzz" 如果 i 同时是 3 和 5 的倍数。
//answer[i] == "Fizz" 如果 i 是 3 的倍数。
//answer[i] == "Buzz" 如果 i 是 5 的倍数。
//answer[i] == i 如果上述条件全不满足。
//
func fizzBuzz(n int) (ans []string) {
	for i := 1; i <= n; i++ {
		var b strings.Builder
		if i%3 == 0 {
			b.WriteString("Fizz")
		}
		if i%5 == 0 {
			b.WriteString("Buzz")
		}
		if b.Len() == 0 {
			b.WriteString(strconv.Itoa(i))
		}
		ans = append(ans, b.String())
	}
	return
}
