package main

import (
	"fmt"
	"strings"
)

//You are given three strings S
//1
//​
// ,S
//2
//​
// ,S
//3
//​
//  consisting of lowercase English letters, and a string T consisting of 1, 2, 3.
//
//Concatenate the three strings according to the characters in T and print the resulting string. Formally, conform to the following instructions.
//
//For each integer i such that 1≤i≤∣T∣, let the string s
func main() {
	var s1, s2, s3, t string
	fmt.Scan(&s1, &s2, &s3, &t)

	var b strings.Builder
	for _, v := range t {
		switch v {
		case '1':
			b.WriteString(s1)
		case '2':
			b.WriteString(s2)
		case '3':
			b.WriteString(s3)
		}
	}
	fmt.Println(b.String())
}
