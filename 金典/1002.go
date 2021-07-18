package coding

// 编写一种方法，对字符串数组进行排序，将所有变位词组合在一起。变位词是指字母相同，但排列不同的字符串。
func groupAnagrams(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, str := range strs {
		var cnt [26]int
		for _, ch := range str {
			cnt[ch-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
