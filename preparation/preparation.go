package preparation

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	Mod1000000007 = 1000000007
	Mod998244353  = 998244353
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

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

var sc, wr = bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)

func scanString() string    { sc.Scan(); return sc.Text() }
func scanRunes() []rune     { return []rune(scanString()) }
func scanInt() int          { a, _ := strconv.Atoi(scanString()); return a }
func scanInt64() int64      { a, _ := strconv.ParseInt(scanString(), 10, 64); return a }
func scanFloat() float64    { a, _ := strconv.ParseFloat(scanString(), 64); return a }
func atoi(s string) int     { x, _ := strconv.Atoi(s); return x }
func atoi64(s string) int64 { x, _ := strconv.ParseInt(s, 10, 64); return x }
func pow(a, b int) int      { return int(math.Pow(float64(a), float64(b))) }

func scanInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	return a
}

func debug(a ...interface{}) {
	if os.Getenv("ONLINE_JUDGE") == "false" {
		fmt.Fprintln(os.Stderr, a...)
	}
}

// 原地去重
func unique(a []int) []int {
	k := 0
	for _, v := range a[1:] {
		if a[k] != v {
			k++
			a[k] = v
		}
	}
	return a[:k+1]
}

func a() {
	//初始化map
	//mp := make(map[int]int)
	//for _, t := range tasks {
	//	mp[t]++
	//}

	//背包问题
	//for _, coin := range coins {
	//	for i := coin; i <= amount; i++ {
	//		dp[i] = min(dp[i], dp[i-coin]+1)
	//	}
	//}
}
