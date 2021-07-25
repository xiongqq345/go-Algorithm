package _51_周赛

func maxCompatibilitySum(students [][]int, mentors [][]int) int {
	m, n := len(students), len(students[0])

	matchS := make([]bool, m)
	matchM := make([]bool, m)

	total := 0
	for s := n; s > 0; s-- {
		for i := range students {
			if matchS[i] {
				continue
			}
			for j := range mentors {
				if matchM[j] {
					continue
				}
				var score int
				for k := range students[0] {
					if students[i][k] == mentors[j][k] {
						score++
					}
				}
				if score == s {
					matchS[i], matchM[j] = true, true
					total += s
				}
			}
		}
	}
	return total
}

[[0,1,0,1,1,1],[1,0,0,1,0,1],[1,0,1,1,0,0]]
[[0,1,0,0,1,1],[0,1,0,0,1,1],[1,0,0,0,0,1]]
5 ,5 ,0