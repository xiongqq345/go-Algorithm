package jz_offer

// 在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。
func firstUniqChar(s string) byte {
	var set [128]int
	for _, v := range s {
		set[v]++
	}
	for _, v := range s {
		if set[v] == 1 {
			return byte(v)
		}
	}
	return ' '
}
