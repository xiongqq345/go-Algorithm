package main

import (
	"bufio"
	"fmt"
)

//We have a grid with H horizontal rows and W vertical columns, where each square contains an integer. The integer written on the square at the i-th row from the top and j-th column from the left is A
//i,j
//
//Determine whether the grid satisfies the condition below.

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
	h, w := scanInt(), scanInt()
	grid := make([][]int, h)
	for i := range grid {
		grid[i] = scanInts(w)
	}
	for i := 0; i < h-1; i++ {
		for j := 0; j < w-1; j++ {
			if grid[i][j]+grid[i+1][j+1] > grid[i+1][j]+grid[i][j+1] {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}
