package greedy

import "sort"

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	var cnt, index int
	for coins > 0 && index < len(costs) {
		if costs[index] <= coins {
			coins -= costs[index]
			cnt++
			index++
		} else {
			break
		}
	}
	return cnt
}
