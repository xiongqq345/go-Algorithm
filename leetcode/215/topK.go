package _15

import (
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	rand.NewSource(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(arr []int, l, r, k int) int {
	mid := partition(arr, l, r)
	if mid == k {
		return arr[mid]
	} else if mid > k {
		return quickSelect(arr, l, mid-1, k)
	} else {
		return quickSelect(arr, mid+1, r, k)
	}
}

func partition(a []int, l, r int) int {
	p := rand.Int()%(r-l+1) + l
	a[p], a[r] = a[r], a[p]

	x, pos := a[r], l
	for i := l; i < r; i++ {
		if a[i] < x {
			a[pos], a[i] = a[i], a[pos]
			pos++
		}
	}
	a[pos], a[r] = a[r], a[pos]
	return pos
}
