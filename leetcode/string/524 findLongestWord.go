package string

import "sort"

// 给你一个字符串 s 和一个字符串数组 dictionary 作为字典，找出并返回字典中最长的字符串，该字符串可以通过删除 s 中的某些字符得到。
//
//如果答案不止一个，返回长度最长且字典序最小的字符串。如果答案不存在，则返回空字符串。
func findLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		if len(dictionary[i]) != len(dictionary[j]) {
			return len(dictionary[i]) > len(dictionary[j])
		}
		return dictionary[i] < dictionary[j]
	})
	for _, w := range dictionary {
		if match(s, w) {
			return w
		}
	}
	return ""
}

func match(s, sub string) bool {
	for i, j := 0, 0; i < len(s); i++ {
		if s[i] == sub[j] {
			j++
		}
		if j == len(sub) {
			return true
		}
	}
	return false
}
