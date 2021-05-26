package binary_search

//搜索旋转数组。给定一个排序后的数组，包含n个整数，但这个数组已被旋转过很多次了，次数不详。请编写代码找出数组中的某个元素，假设数组元素原先是按升序排列的。若有多个相同元素，返回索引值最小的一个。

func search(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l < r {
		mid := (l + r) / 2

		if arr[mid] > arr[l] {

		} else if arr[mid] < arr[l] {

		}
	}
	return -1
}
