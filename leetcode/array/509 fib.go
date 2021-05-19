package array

func fib(n int) int {
	if n < 2 {
		return n
	}
	pre, cur := 0, 1
	for i := 2; i <= n; i++ {
		pre, cur = cur, pre+cur
	}
	return cur
}

//func fib(n int) int {
//	sqrt5 := math.Sqrt(5)
//	p1 := math.Pow((1+sqrt5)/2, float64(n))
//	p2 := math.Pow((1-sqrt5)/2, float64(n))
//	return int(math.Round((p1 - p2) / sqrt5))
//}
