package 力扣杯竞赛

import (
	"math"
	"strconv"
)

const mod int = 1e9 + 7

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func atoi(s string) int     { x, _ := strconv.Atoi(s); return x }
func atoi64(s string) int64 { x, _ := strconv.ParseInt(s, 10, 64); return x }
func pow(a, b int) int      { return int(math.Pow(float64(a), float64(b))) }
func pow2(a int) int        { return int(math.Pow(2, float64(a))) }

func AMax(p *int, xs ...int) { *p = Max(*p, Max(xs...)) }
func AMin(p *int, xs ...int) { *p = Min(*p, Min(xs...)) }
func Max(xs ...int) int {
	max := xs[0]
	for _, x := range xs {
		if max < x {
			max = x
		}
	}
	return max
}

func Min(xs ...int) int {
	min := xs[0]
	for _, x := range xs {
		if min > x {
			min = x
		}
	}
	return min
}

type heapInt []int

//Less  小于就是小根堆，大于号就是大根堆
func (h *heapInt) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *heapInt) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *heapInt) Len() int           { return len(*h) }
func (h *heapInt) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *heapInt) Pop() interface{}   { t := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return t }
func (h *heapInt) Peek() int          { return (*h)[0] }
