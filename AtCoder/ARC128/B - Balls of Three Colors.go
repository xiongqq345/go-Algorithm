package b

//We have R red balls, G green balls, and B blue balls. You can do the following operation any number of times:
//
//choose two balls of different colors and turn them into two balls of the remaining color.
//For example, you can choose a red ball and a blue ball and turn them into two green balls.
//
//Your objective is to make all balls have the same color. Determine whether this objective is achievable. If it is, find the minimum number of operations required to achieve it.
//
//For each input file, solve T test cases.

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc, wr = bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
)

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

func main() {
	sc.Split(bufio.ScanWords)
	t := scanInt()
	for i := 0; i < t; i++ {
		f(scanInts(3))
	}
}

func f(a []int) {

	fmt.Println()
}

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
