package _35

func searchInsert(nums []int, target int) int {
	n := len(nums)
	if nums[n-1] < target {
		return n
	}

	l, r := 0, n-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
