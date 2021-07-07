package binary_search

// 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
func search(arr []int, target int) int {
	i, j := 0, len(arr)-1
	for i <= j {
		h := int(uint(i+j) >> 1)
		if arr[h] < target {
			i = h + 1
		} else if arr[h] > target {
			j = h - 1
		} else {
			return h
		}
	}
	return -1
}
