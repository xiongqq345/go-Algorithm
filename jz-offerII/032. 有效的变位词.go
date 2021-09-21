package jz_offerII

//给定两个字符串 s 和 t ，编写一个函数来判断它们是不是一组变位词（字母异位词）。
//
//注意：若 s 和 t 中每个字符出现的次数都相同且字符顺序不完全相同，则称 s 和 t 互为变位词（字母异位词）。
//
//
func isAnagram(s string, t string) bool {
	if len(s) != len(t) || s == t {
		return false
	}
	var a1, a2 [26]int
	for i := range s {
		a1[s[i]-'a']++
		a2[t[i]-'a']++
	}
	return a1 == a2
}
