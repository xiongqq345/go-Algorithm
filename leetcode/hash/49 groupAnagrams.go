package hash

// 给定一个字符串数组，将字母异位词组合在一起。可以按任意顺序返回结果列表。
//
//字母异位词指字母相同，但排列不同的字符串。
func groupAnagrams(strs []string) [][]string {
	var ans [][]string
	mp := make(map[[26]int]int)
	for _, str := range strs {
		var tmp [26]int
		for _, ch := range str {
			tmp[ch-'a']++
		}

		if pos, ok := mp[tmp]; ok {
			ans[pos] = append(ans[pos], str)
		} else {
			ans = append(ans, []string{})
			mp[tmp] = len(ans) - 1
			ans[len(ans)-1] = append(ans[len(ans)-1], str)
		}
	}
	return ans
}
