package main

import (
	"fmt"
	"sort"
)

// 链接：https://ac.nowcoder.com/acm/contest/18713/G
//来源：牛客网
//
//题目描述
//Shenyang's night fair culture is developed very well. Every time Bob comes to Shenyang, he will definitely go to a night fair called {The Witchwood}TheWitchwood. There are {n}n snack stalls in {The Witchwood}TheWitchwood, the {i}ith of which gives him a_ia
//i
//
//  pleasure.
//Bob's stomach allows him to eat {k}k snack stalls at most. So Bob wants to know the maximum pleasure he can get after visiting the night market.

//输入描述:
//Print one integer denoting the maximum pleasure Bob can get.

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	sort.Ints(a)
	var ans int
	for i := n - 1; i >= 0 && k > 0; i-- {
		ans += a[i]
		k--
	}
	fmt.Println(ans)
}
