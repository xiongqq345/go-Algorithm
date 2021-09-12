package AtCoder

import (
	"bufio"
	"fmt"
	"os"
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

type heapInt []int

//Less  小于就是小根堆，大于号就是大根堆
func (h *heapInt) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *heapInt) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *heapInt) Len() int           { return len(*h) }
func (h *heapInt) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *heapInt) Pop() interface{} {
	t := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return t
}
func (h *heapInt) Peek() int {
	return (*h)[0]
}

var sc, wr = bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)

func scanString() string { sc.Scan(); return sc.Text() }
func scanRunes() []rune  { return []rune(scanString()) }
func scanInt() int       { a, _ := strconv.Atoi(scanString()); return a }
func scanInt64() int64   { a, _ := strconv.ParseInt(scanString(), 10, 64); return a }
func scanFloat() float64 { a, _ := strconv.ParseFloat(scanString(), 64); return a }

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
