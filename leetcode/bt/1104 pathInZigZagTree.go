package bt

import "math"

func pathInZigZagTree(label int) []int {
	var getLevelMinLabel func(int) int
	getLevelMinLabel = func(i int) int {
		return int(math.Pow(2, float64(i)))
	}
	level := int(math.Log2(float64(label)))
	ans := make([]int, level+1)
	ans[level] = label
	level--
	for level >= 0 {
		sum := getLevelMinLabel(level) + getLevelMinLabel(level+1) - 1
		label /= 2
		label = sum - label
		ans[level] = label
		level--
	}

	return ans
}
