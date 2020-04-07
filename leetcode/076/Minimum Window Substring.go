package leetcode

func minWindow(s string, t string) string {

	if len(s) < len(t) {
		return ""
	}
	hash := make([]int, 256)
	for i := 0; i < len(t); i++ {
		hash[t[i]]++
	}

	l, r := 0, 0
	sw := s[l:r]
	minStr := sw
	for r = 1; r < len(s); r++ {
		if hash[s[r]] == 1 {
			hash[s[r]]++
		}
		switch hash[s[l]] {
		case 0:
			l++
		case 2:
		case 3:
			l++
			hash[s[l]]--
		default:
		}
		c := 0
		for c < len(t) {
			if hash[t[c]] < 2 {
				break
			}
			c++
		}
		if c == len(t) {
			if len(minStr) == 0 {
				minStr = sw
			} else if len(sw) < len(minStr) {
				minStr = sw
			}
		}
		sw = s[l:r]

	}
	return minStr
}
