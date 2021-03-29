package _61

func hammingDistance(x int, y int) int {
	var res int
	z := x ^ y
	for z > 0 {
		res += z & 1
		z >>= 1
	}
	return res
}
