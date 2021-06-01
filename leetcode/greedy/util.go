package greedy

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
