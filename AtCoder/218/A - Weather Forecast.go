package main

import "fmt"

//You are given a string S, which represents a weather forecast for the seven days starting tomorrow.
//The i-th of those seven days is forecast to be sunny if the i-th character of S is o, and rainy if that character is x.
//
//Tell us whether the N-th day is forecast to be sunny.

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)
	if s[n-1] == 'x' {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
