package math

//给定一个正整数 n ，你可以做如下操作：
//
//如果 n 是偶数，则用 n / 2替换 n 。
//如果 n 是奇数，则可以用 n + 1或n - 1替换 n 。
//n 变为 1 所需的最小替换次数是多少？
//
func integerReplacement(n int) int {
	var ans int
	for n != 1 {
		switch {
		case n%2 == 0:
			ans--
			n /= 2
		case n%4 == 1:
			n /= 2
		case n == 3:
			n = 1
		default:
			n = n/2 + 1
		}
		ans += 2
	}
	return ans
}
