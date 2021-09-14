package jz_offerII

//给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的某个变位词。
//
//换句话说，第一个字符串的排列之一是第二个字符串的 子串 。
func checkInclusion(s1 string, s2 string) bool {
	var s, t [26]int
	for _, ch := range s1 {
		s[ch-'a']++
	}
	for i := range s2 {
		t[s2[i]-'a']++
		if i >= len(s1) {
			t[s2[i-len(s1)]-'a']--
		}
		if s == t {
			return true
		}
	}
	return false
}
