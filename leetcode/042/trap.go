package leetcode

func trap(height []int) int {
	lens := len(height)
	if lens <= 1 {
		return 0
	}
	l, r, lh, rh, sum := 0, lens-1, 0, 0, 0
	for l < r {
		if height[l] < height[r] {
			if height[l] < lh {
				sum += lh - height[l]
			} else {
				lh = height[l]
			}
			l++
		} else {
			if height[r] < rh {
				sum += rh - height[r]
			} else {
				rh = height[r]
			}
			r--
		}
	}
	return sum
}
