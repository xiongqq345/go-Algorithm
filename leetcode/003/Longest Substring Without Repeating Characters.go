package leetcode

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	maxLen :=1
	left := 0
	right := 1
	str := s[:1]
	for i := 1; i < len(s); i++ {
		if index := strings.IndexByte(str, s[i]); index != -1 {
			left += index + 1
		}
		right++
		str = s[left:right]
		if len(str)>maxLen {
			maxLen = len(str)
		}
	}
	return maxLen
}
