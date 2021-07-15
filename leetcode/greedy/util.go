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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
