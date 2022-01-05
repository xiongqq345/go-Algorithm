package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Mr. Chef is the manager of the Code cinemas and after a long break, the theatres are now open to the public again. To compensate for the loss in revenue due to Covid-19, Mr. Chef wants to maximize the profits for every show from now on and at the same time follow the guidelines set the by government. The guidelines are:
//
//If two people are seated in the same row, there must be at least one empty seat between them.
//If two people are seated in different rows, there must be at least one completely empty row between them. That is, if there are people seated in rows i and j where i<j, there must be some row k such that i<k<j and nobody is seated in row k.
//Given the information about the number of rows and the number of seats in each row, find the maximum number of tickets Mr. Chef can sell.
//
//Input Format
//The first line of input will contain a single integer T, denoting the number of test cases. The description of T test cases follows.
//Each test case consists of a single line of input containing two space-separated integers N,M — the number of rows and the number of seats in each row, respectively.
//Output Format
//For each test case, output a single line containing one integer – the maximum number of tickets Mr. Chef can sell.

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
		a, b := (scanInt()+1)/2, (scanInt()+1)/2
		fmt.Println(a * b)
	}
}
