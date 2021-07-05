package _48_周赛

func eliminateMaximum(dist []int, speed []int) int {
	n := len(dist)
	mp := make(map[int]int)
	var maxRound int
	for i := 0; i < n; i++ {
		var round int
		if dist[i]%speed[i] > 0 {
			round = dist[i]/speed[i] + 1
		} else {
			round = dist[i] / speed[i]
		}
		mp[round]++
		maxRound = max(maxRound, round)
	}
	bucket := make([]int, maxRound+1)
	for round, v := range mp {
		bucket[round] = v
	}

	var kill, lastIndex int
	for ; kill < n; kill++ {
		for i := lastIndex; i <= maxRound; i++ {
			if bucket[i] > 0 {
				if i <= kill {
					return kill
				}

				lastIndex = i
				bucket[i]--
				break
			}
		}
	}
	return kill
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
