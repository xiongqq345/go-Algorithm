package jz_offer

// 在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。

func reversePairs(nums []int) int {
	var count int
	var mergeSort func([]int) []int
	mergeSort = func(arr []int) []int {
		if len(arr) < 2 {
			return arr
		}
		middle := len(arr) / 2
		l := mergeSort(arr[0:middle])
		r := mergeSort(arr[middle:])
		return func(l, r []int) []int {
			var result []int
			m, n := 0, 0
			for m < len(l) && n < len(r) {
				if l[m] <= r[n] {
					result = append(result, l[m])
					m++
				} else {
					result = append(result, r[n])
					count += len(l) - m
					n++
				}
			}
			result = append(result, r[n:]...)
			result = append(result, l[m:]...)
			return result
		}(l, r)
	}
	mergeSort(nums)
	return count
}
