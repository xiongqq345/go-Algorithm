package main

import (
	"fmt"
	"sort"
)

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(permuteUnique([]byte(s)))
}

func permuteUnique(str []byte) int {
	sort.Slice(str, func(i, j int) bool {
		return str[i] < str[j]
	})
	n := len(str)
	var ans int
	vis := make([]bool, n)
	var helper func(int)
	helper = func(length int) {
		if length == n {
			ans++
			return
		}

		for i := 0; i < n; i++ {
			if i > 0 && str[i-1] == str[i] && vis[i-1] {
				continue
			}
			if vis[i] {
				continue
			}

			vis[i] = true
			helper(length + 1)
			vis[i] = false
		}
	}
	helper(0)
	return ans
}
