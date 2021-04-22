package _81

import "math"

func numRabbits(answers []int) int {
	m := make(map[float64]float64)
	for _, answer := range answers {
		m[float64(answer)]++
	}
	var res float64
	for ans, num := range m {
		if num <= ans {
			res += ans + 1
		} else {
			res += math.Ceil(num/(ans+1)) * (ans + 1)
		}
	}
	return int(res)
}
