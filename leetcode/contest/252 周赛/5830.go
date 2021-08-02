package _52_周赛

// 给你一个整数 n 。如果 n 恰好有三个正除数 ，返回 true ；否则，返回 false 。
//
//如果存在整数 k ，满足 n = k * m ，那么整数 m 就是 n 的一个 除数 。
//
func isThree(n int) bool {
	cnt := 2
	for i := 2; i <= n/2; i++ {
		if n/i*i == n {
			cnt++
		}
	}
	return cnt == 3
}
