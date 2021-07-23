package array

// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
func trap(arr []int) int {
	l, r, lm, rm, ans := 0, len(arr)-1, 0, 0, 0
	for l < r {
		if arr[l] < arr[r] {
			lm = max(lm, arr[l])
			ans += lm - arr[l]
			l++
		} else {
			rm = max(rm, arr[r])
			ans += rm - arr[r]
			r--
		}
	}
	return ans
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
