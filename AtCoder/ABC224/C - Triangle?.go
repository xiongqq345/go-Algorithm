package main

//In the xy-plane, we have N points numbered 1 through N.
//Point i is at the coordinates (X
//i
//​
// ,Y
//i
//​
// ). Any two different points are at different positions.
//Find the number of ways to choose three of these N points so that connecting the chosen points with segments results in a triangle with a positive area.
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func main() {
	sc.Split(bufio.ScanWords)
	n := scanInt()
	ps := make([][2]int, n)
	for i := 0; i < n; i++ {
		ps[i][0], ps[i][1] = scanInt(), scanInt()
	}

	ans := n * (n - 1) * (n - 2) / 6
	var cnt int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if ps[i][0] == ps[j][0] && ps[j][0] == ps[k][0] {
					cnt++
					continue
				}
				if ps[i][0] == ps[j][0] || ps[j][0] == ps[k][0] {
					continue
				}
				if float64(ps[k][1]-ps[j][1])/float64(ps[k][0]-ps[j][0]) == float64(ps[j][1]-ps[i][1])/float64(ps[j][0]-ps[i][0]) {
					cnt++
				}
			}
		}
	}
	fmt.Println(ans - cnt)
}
