package jz_offer

func fib(n int) int {
	pre, cur := 0, 1
	for ; n > 0; n-- {
		pre, cur = cur, (pre+cur)%(1e9+7)
	}
	return pre
}
