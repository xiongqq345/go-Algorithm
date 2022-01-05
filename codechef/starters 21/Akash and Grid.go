package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Akash is stuck in a N×N grid, where N is odd. The rows of the grid are numbered 1 to N from top to bottom, and the columns are numbered 1 to N from left to right. The cell at the intersection of the i-th row and j-th column will be denoted (i,j).
//
//The grid has a unique center cell — ((N+1)/2,(N+1)/2). For example, when N=5 the center is cell (3,3).
//
//Akash is currently at cell (xs,ys). He would like to reach the exit of the grid, which is located at the center. It is guaranteed that (xs,ys) is not the center.
//
//Suppose Akash is at cell (x,y). He can make the following movements:
//
//He can freely move along diagonals, i.e, to cells (x−1,y−1),(x−1,y+1),(x+1,y−1),(x+1,y+1)
//He can move one step horizontally or vertically with the cost of 1 gold coin, i.e, to cells (x,y−1),(x,y+1),(x−1,y),(x+1,y)
//Note that Akash is not allowed to make a move that will take him out of bounds of the grid.
//
//Akash would like to minimize the number of coins required to reach the center. Please help him find this number.
//
//Input Format
//The first line of input contains a single integer T, denoting the number of test cases. The description of T test cases follows.
//Each test case consists of a single line of input, containing three space-separated integers N,xs,ys — the size of the grid and the coordinates of Akash's starting cell.
//Output Format
//For each test case, output in a single line the minimum number of gold coins Akash needs to reach the center.

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
	t := scanInt()
	for i := 0; i < t; i++ {
		_, x, y := scanInt(), scanInt(), scanInt()
		if (x+y)%2 == 0 {
			fmt.Println(0)
		} else {
			fmt.Println(1)
		}
	}
}
