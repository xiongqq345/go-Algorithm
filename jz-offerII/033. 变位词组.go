package jz_offerII

//给定一个字符串数组 strs ，将 变位词 组合在一起。 可以按任意顺序返回结果列表。
//
//注意：若两个字符串中每个字符出现的次数都相同，则称它们互为变位词。
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
