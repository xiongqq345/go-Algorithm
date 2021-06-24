package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var a, b, n uint
	fmt.Scan(&n)
	for ; n > 0; n-- {
		fmt.Scan(&a, &b)
		var maxx int
		for i := a; i <= b; i++ {
			maxx = max(bits.OnesCount(i), maxx)
		}
		fmt.Println(maxx)
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
