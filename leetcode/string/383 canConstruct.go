package string

func canConstruct(ransomNote string, magazine string) bool {
	mp := make(map[int32]int)
	for _, v := range ransomNote {
		mp[v]++
	}
	for _, v := range magazine {
		if mp[v] > 1 {
			mp[v]--
		} else if mp[v] == 1 {
			delete(mp, v)
		}
		if len(mp) == 0 {
			return true
		}
	}
	return false
}
