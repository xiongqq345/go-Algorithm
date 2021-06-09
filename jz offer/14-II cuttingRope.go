package jz_offer

func cuttingRope2(n int) int {
	if n <= 3 {
		return n - 1
	}
	b, p := n%3, 1e9+7
	ans := 1
	for i := 1; i < n/3; i++ {
		ans = ans * 3 % int(p)
	}
	switch b {
	case 1:
		ans = ans * 4 % int(p)
	case 2:
		ans = ans * 6 % int(p)
	case 0:
		ans = ans * 3 % int(p)
	}
	return ans
}
