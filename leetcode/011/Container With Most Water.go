package leetcode

func maxArea(height []int) int {
	max, high, l, r := 0, 0, 0, len(height)-1
	for l != r {
		left, right := height[l], height[r]
		if left < right {
			high = left
			l++
		} else {
			high = right
			r--
		}
		tmp := (r - l + 1) * high
		if max < tmp {
			max = tmp
		}
	}
	return max
}
