package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

//Snuke has 1 gram of gold and 0 grams of silver now. He will do trading of gold and silver for the following N days. On each day, he has two choices: do nothing, or make a trade. If he trades on Day i (1≤i≤N), the following will happen.
//
//If he has x grams of gold before the trade, exchange all of it for x×A
//i
//​
//  grams of silver. On the other hand, if he has x grams of silver, exchange all of it for x/A
//i
//​
//  grams of gold.
//Snuke's objective is to maximize the amount of gold he has in the end. Find one way to achieve his objective.
func main() {
	sc.Split(bufio.ScanWords)
	n := scanInt()
	a := scanInts(n)
	g := true
	op := make([]string, n)
	for i := range op {
		op[i] = "0"
	}

	for i := 0; i < n-1; i++ {
		if g {
			if a[i] > a[i+1] {
				g = false
				op[i] = "1"
			}
		} else {
			if a[i] <= a[i+1] {
				g = true
				op[i] = "1"
			}
		}
	}
	if a[n-2] > a[n-1] {
		op[n-1] = "1"
	}

	fmt.Println(strings.Join(op, " "))
}
