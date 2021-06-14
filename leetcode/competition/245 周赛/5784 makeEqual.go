package _45_周赛

func makeEqual(words []string) bool {
	mp := make(map[int32]int)
	for _, str := range words {
		for _, c := range str {
			mp[c]++
		}
	}
	n := len(words)
	for _, v := range mp {
		if v%n != 0 {
			return false
		}
	}
	return true
}
