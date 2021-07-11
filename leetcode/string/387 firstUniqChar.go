package string

// 给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
func firstUniqChar(s string) int {
	var cnt [26]int
	for _, v := range s {
		cnt[v-'a']++
	}
	for i, v := range s {
		if cnt[v-'a'] == 1 {
			return i
		}
	}
	return -1
}
