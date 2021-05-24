package math

func numRabbits(answers []int) int {
	m, res := make(map[int]int), 0
	for _, v := range answers {
		m[v]++
	}
	for answer, num := range m {
		res += (num + answer) / (answer + 1) * (answer + 1)
	}
	return res
}
