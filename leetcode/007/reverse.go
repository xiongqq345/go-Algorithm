package _07

import "math"

func reverse(x int) int {
	var rev int
	for x != 0 {
		if rev > math.MaxInt32/10 || rev < math.MinInt32/10 {
			return 0
		}
		rev = rev*10 + x%10
		x /= 10

	}
	return rev
}
