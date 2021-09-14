package jz_offerII

//给定两个字符串 s 和 p，找到 s 中所有 p 的 变位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
//
//变位词 指字母相同，但排列不同的字符串。
//
func findAnagrams(s string, p string) []int {
	var ans []int
	var set [26]int
	for _, ch := range p {
		set[ch-97]++
	}
	var tmp [26]int
	for i := range s {
		tmp[s[i]-97]++
		if i >= len(p) {
			tmp[s[i-len(p)]-97]--
		}
		if set == tmp {
			ans = append(ans, i-len(p)+1)
		}
	}
	return ans
}
