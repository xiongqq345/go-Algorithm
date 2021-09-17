package jz_offerII

import (
	"math/rand"
	"time"
)

// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//
//请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
func findKthLargest(nums []int, k int) int {
	rand.NewSource(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(a []int, l, r, k int) int {
	q := partition(a, l, r)
	if q == k {
		return a[q]
	} else if q > k {
		return quickSelect(a, l, q-1, k)
	} else {
		return quickSelect(a, q+1, r, k)
	}
}

func partition(a []int, l, r int) int {
	p := rand.Int()%(r-l+1) + l
	a[r], a[p] = a[p], a[r]

	p = l
	for i := l; i < r; i++ {
		if a[i] < a[r] {
			a[p], a[i] = a[i], a[p]
			p++
		}
	}
	a[r], a[p] = a[p], a[r]
	return p
}
