package coding

// 搜索旋转数组。给定一个排序后的数组，包含n个整数，但这个数组已被旋转过很多次了，次数不详。请编写代码找出数组中的某个元素，假设数组元素原先是按升序排列的。若有多个相同元素，返回索引值最小的一个。
//
func search(arr []int, target int) int {
	i, j := 0, len(arr)-1
	for i < j {
		h := int(uint(i+j) >> 1)
		if arr[i] == target {
			return i
		}
		if arr[h] < arr[i] {
			if target < arr[h] || target > arr[j] {
				j = h
			} else {
				i = i + 1
			}
		} else if arr[h] > arr[i] {
			if target > arr[h] || target < arr[i] {
				i = h + 1
			} else {
				j = h
			}
		} else {
			i++
		}
	}
	if arr[i] != target {
		return -1
	}
	return i
}
