package binary_search

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l, r := 0, len(nums)-1
	start, end := 0, len(nums)-1
	for l+1 < r {
		mid := l + (r-l)/2
		if nums[mid] == target && nums[mid] > nums[mid-1] {
			start = mid
			break
		}
		if nums[mid] < target {
			l = mid
		} else {
			r = mid
		}
	}
	r = len(nums) - 1
	for l+1 < r {
		mid := l + (r-l)/2
		if nums[mid] == target && nums[mid] < nums[mid+1] {
			end = mid
			break
		}
		if nums[mid] < target+1 {
			l = mid
		} else {
			r = mid
		}
	}

	if nums[start] == target && nums[end] == target {
		return []int{start, end}
	}
	if nums[start] == target {
		return []int{start, start}
	}
	if nums[end] == target {
		return []int{end, end}
	}
	return []int{-1, -1}
}
