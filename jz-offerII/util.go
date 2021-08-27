package jz_offerII

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
