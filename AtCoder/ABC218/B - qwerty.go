package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// You are given a sequence of 26 integers P=(P
//1
//​
// ,P
//2
//​
// ,…,P
//26
//​
// ) consisting of integers from 1 through 26. It is guaranteed that all elements in P are distinct.
//
//Print a string S of length 26 that satisfies the following condition.
//
//For every i (1≤i≤26), the i-th character of S is the lowercase English letter that comes P
//i
//​
// -th in alphabetical order.

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	arr := strings.Split(sc.Text(), " ")
	for _, n := range arr {
		num, _ := strconv.Atoi(n)
		fmt.Printf("%c", num+96)
	}
	fmt.Println()
}
