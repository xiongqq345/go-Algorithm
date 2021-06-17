package string

import "strings"

func lengthOfLastWord(s string) int {
	words := strings.Fields(s)
	if len(words) > 0 {
		return len(words[len(words)-1])
	}
	return 0
}
