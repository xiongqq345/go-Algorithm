package math

// 给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
func findMaxLength(nums []int) int {
	m := map[int]int{0: -1}
	var sum, length int
	for i, v := range nums {
		if v == 0 {
			sum--
		} else {
			sum++
		}

		if val, ok := m[sum]; !ok {
			m[sum] = i
		} else {
			length = max(length, i-val)
		}
	}
	return length
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
