package sort

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	l := MergeSort(arr[:mid])
	r := MergeSort(arr[mid:])
	return merge(l, r)
}

func merge(l, r []int) []int {
	var result []int
	var m, n int
	for m < len(l) && n < len(r) {
		if l[m] <= r[n] {
			result = append(result, l[m])
			m++
		} else {
			result = append(result, r[n])
			n++
		}
	}
	result = append(result, r[n:]...)
	result = append(result, l[m:]...)
	return result
}
