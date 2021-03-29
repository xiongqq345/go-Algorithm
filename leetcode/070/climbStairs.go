package _70

func climbStairs2(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 2
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairs(n int) int {
	p, q, r := 1, 2, 0
	for i := 3; i <= n; i++ {
		r = p + q
		p, q = q, r
	}
	return r
}
