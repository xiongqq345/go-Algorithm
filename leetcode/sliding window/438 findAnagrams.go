package sliding_window

// 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
//
//异位词 指字母相同，但排列不同的字符串。
//
func findAnagrams(s string, p string) []int {
	var ans []int
	var arr [26]int
	for _, v := range p {
		arr[v-'a']++
	}

	var tmp [26]int
	for i, j := 0, 0; j < len(s); j++ {
		tmp[s[j]-'a']++
		if tmp == arr {
			ans = append(ans, i)
		}
		if j-i >= len(p)-1 {
			tmp[s[i]-'a']--
			i++
		}
	}
	return ans
}
