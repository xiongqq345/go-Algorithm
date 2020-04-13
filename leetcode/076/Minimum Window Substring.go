package leetcode

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	hash := make([]int, 256)
	for i := 0; i < len(t); i++ {
		hash[t[i]]++
	}

	l, count, min, res := 0, len(t), len(s)+1, ""
	for r := 0; r < len(s); r++ {
		hash[s[r]]--
		if hash[s[r]] >= 0 {
			count--
		}
		for l < r && hash[s[l]] < 0 {
			hash[s[l]]++
			l++
		}
		if count == 0 && min > r-l+1 {
			min = r - l + 1
			res = s[l : r+1]
		}
	}

	return res

}
