package _5_双周赛

import (
	"strings"
)

func removeOccurrences(s string, part string) string {
	var pre string
	for pre != s {
		pre = s
		s = strings.Replace(s, part, "", 1)
	}
	return s
}
