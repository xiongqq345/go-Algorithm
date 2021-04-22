package sort

func QuickSort(arr []int, l, r int) {
	if l < r {
		mid := partition(arr, l, r)
		QuickSort(arr, l, mid-1)
		QuickSort(arr, mid+1, r)
	}
}

func partition(arr []int, l, r int) int {
	key, loc := arr[r], l
	for i := l; i < r; i++ {
		if arr[i] < key {
			arr[loc], arr[i] = arr[i], arr[loc]
			loc++
		}
	}
	arr[loc], arr[r] = arr[r], arr[loc]
	return loc
}
