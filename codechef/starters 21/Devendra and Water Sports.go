package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Recently, Devendra went to Goa with his friends. Devendra is well known for not following a budget.
//
//He had Rs. Z at the start of the trip and has already spent Rs. Y on the trip. There are three kinds of water sports one can enjoy, with prices Rs. A, B, and C. He wants to try each sport at least once.
//
//If he can try all of them at least once output YES, otherwise output NO.
//
//Input Format
//The first line of input contains a single integer T, denoting the number of test cases. The description of T test cases follows.
//Each test case consists of a single line of input containing five space-separated integers Z,Y,A,B,C.
//Output Format
//For each test case, print a single line containing the answer to that test case — YES if Devendra can try each ride at least once, and NO otherwise.
//
//You may print each character of the string in uppercase or lowercase (for example, the strings "yEs", "yes", "Yes" and "YES" will all be treated as identical).

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
		a := scanInts(5)
		if a[0] >= a[1]+a[2]+a[3]+a[4] {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
