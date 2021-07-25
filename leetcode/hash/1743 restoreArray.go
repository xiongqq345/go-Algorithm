package hash

func restoreArray(adjacentPairs [][]int) []int {
	n := len(adjacentPairs) + 1
	mp := make(map[int][]int, n)
	for _, v := range adjacentPairs {
		mp[v[0]] = append(mp[v[0]], v[1])
		mp[v[1]] = append(mp[v[1]], v[0])
	}
	var ans []int
	var element int
	for k, v := range mp {
		if len(v) == 1 {
			element = v[0]
			ans = append(ans, k, v[0])
			break
		}
	}
	for len(mp[element]) == 2 {
		v1, v2 := mp[element][0], mp[element][1]
		if v1 == ans[len(ans)-2] {
			element = v2
		} else {
			element = v1
		}
		ans = append(ans, element)
	}
	return ans
}
