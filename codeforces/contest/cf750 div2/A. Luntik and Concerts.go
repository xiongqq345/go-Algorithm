package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Luntik has decided to try singing. He has 𝑎 one-minute songs, 𝑏 two-minute songs and 𝑐 three-minute songs. He wants to distribute all songs into two concerts such that every song should be included to exactly one concert.
//
//He wants to make the absolute difference of durations of the concerts as small as possible. The duration of the concert is the sum of durations of all songs in that concert.
//
//Please help Luntik and find the minimal possible difference in minutes between the concerts durations.

//
var sc, wr = bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)

func scanString() string { sc.Scan(); return sc.Text() }
func scanRunes() []rune  { return []rune(scanString()) }
func scanInt() int       { a, _ := strconv.Atoi(scanString()); return a }
func scanInt64() int64   { a, _ := strconv.ParseInt(scanString(), 10, 64); return a }
func scanFloat() float64 { a, _ := strconv.ParseFloat(scanString(), 64); return a }

func scanInt64s(n int) []int64 {
	a := make([]int64, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt64()
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)
	n := scanInt()
	for i := 0; i < n; i++ {
		a := scanInt64s(3)
		fmt.Println((a[0] + 2*a[1] + 3*a[2]) % 2)
	}
}
