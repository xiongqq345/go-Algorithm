package string

// 给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
//
//换句话说，第一个字符串的排列之一是第二个字符串的 子串 。
func checkInclusion(s1 string, s2 string) bool {
	m := len(s1)
	var cnt1, cnt2 [26]int
	for _, v := range s1 {
		cnt1[v-'a']++
	}

	for i := range s2 {
		cnt2[s2[i]-'a']++
		if i >= m {
			cnt2[s2[i-m]-'a']--
		}
		if cnt1 == cnt2 {
			return true
		}
	}
	return false
}
