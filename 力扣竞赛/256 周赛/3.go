package _56_周赛

import "sort"

func minSessions(tasks []int, sessionTime int) int {
	mp := make(map[int]int)
	for _, t := range tasks {
		mp[t]++
	}

	sort.Ints(tasks)
	var cnt int
	for i := len(tasks) - 1; i >= 0; i-- {
		if mp[tasks[i]] <= 0 {
			continue
		}
		cnt++
		t := sessionTime - tasks[i]
		for need := t; t > 0 && need > 0; {
			if mp[need] > 0 {
				mp[need]--
				t -= need
				need = t
			} else {
				need--
			}
		}
	}
	return cnt
}
