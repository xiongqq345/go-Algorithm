package string

// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var cnt [26]int
	for i := range s {
		cnt[s[i]-'a']++
		cnt[t[i]-'a']--
	}
	for _, v := range cnt {
		if v != 0 {
			return false
		}
	}
	return true
}
