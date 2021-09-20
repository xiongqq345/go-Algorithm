package main

import (
	"fmt"
	"strings"
)

//A bracket sequence is a string containing only characters "(" and ")". A regular bracket sequence is a bracket sequence that can be transformed into a correct arithmetic expression by inserting characters "1" and "+" between the original characters of the sequence. For example, bracket sequences "()()" and "(())" are regular (the resulting expressions are: "(1)+(1)" and "((1+1)+1)"), and ")(", "(" and ")" are not.
//
//You are given an integer 𝑛. Your goal is to construct and print exactly 𝑛 different regular bracket sequences of length 2𝑛.
//
//Input
//The first line contains one integer 𝑡 (1≤𝑡≤50) — the number of test cases.
//
//Each test case consists of one line containing one integer 𝑛 (1≤𝑛≤50).
//
//Output
//For each test case, print 𝑛 lines, each containing a regular bracket sequence of length exactly 2𝑛. All bracket sequences you output for a testcase should be different (though they may repeat in different test cases). If there are multiple answers, print any of them. It can be shown that it's always possible.

func main() {
	var t, n int
	mp := make(map[int][]string)
	for fmt.Scan(&t); t > 0; t-- {
		fmt.Scan(&n)
		if arr, ok := mp[n]; ok {
			for _, v := range arr {
				fmt.Println(v)
			}
		} else {
			ans := generateParenthesis(n)
			mp[n] = ans
			for _, v := range ans {
				fmt.Println(v)
			}
		}
	}
}

func generateParenthesis(n int) []string {
	s := []byte(strings.Repeat("(", n) + strings.Repeat(")", n))
	ans := []string{string(s)}
	for i := 1; i < n; i++ {
		s[i], s[2*n-1-i] = s[2*n-1-i], s[i]
		ans = append(ans, string(s))
		s[i], s[2*n-1-i] = s[2*n-1-i], s[i]
	}
	return ans
}
