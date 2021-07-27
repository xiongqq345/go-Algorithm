package string

// 给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
//
//在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。
//
//注意:
//假设字符串的长度不会超过 1010。
//
func longestPalindrome(s string) int {
	mp := make(map[rune]int)
	for _, v := range s {
		mp[v]++
	}

	var hasOne bool
	var ans int
	for _, v := range mp {
		if v%2 == 1 {
			hasOne = true
		}
		ans += v / 2 * 2
	}
	if hasOne {
		ans++
	}
	return ans
}
