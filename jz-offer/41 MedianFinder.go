package jz_offer

import "sort"

type MedianFinder struct {
	data []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}

func (mf *MedianFinder) AddNum(num int) {
	index := sort.SearchInts(mf.data, num)
	mf.data = append(mf.data[:index], append([]int{num}, mf.data[index:]...)...)
}

func (mf *MedianFinder) FindMedian() float64 {
	n := len(mf.data)
	if n == 0 {
		return 0
	}
	if n%2 == 1 {
		return float64(mf.data[n/2])
	}
	return float64(mf.data[n/2]+mf.data[n/2-1]) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
