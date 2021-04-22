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

func partition(arr []int, l, r int) int {
	p := rand.Int()%(r-l+1) + l
	arr[p], arr[r] = arr[r], arr[p]

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
