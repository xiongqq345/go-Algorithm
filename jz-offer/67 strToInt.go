package jz_offer

import (
	"math"
	"strings"
)

func strToInt(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	sign := 1
	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			sign = -1
		}
		s = s[1:]
	}
	var n int
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			n = 10*n + int(ch-'0')
		} else {
			break
		}
		if n > math.MaxInt32 {
			if sign == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}

	return sign * n
}
