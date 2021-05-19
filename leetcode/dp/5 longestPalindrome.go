package dp

// 最长回文子串
func longestPalindrome(s string) string {
	var res string
	for i := range s {
		s1 := Palindrome(s, i, i)
		s2 := Palindrome(s, i, i+1)
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return res
}

func Palindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}

func longestPalindrome2(s string) string {
	center := 0
	maxLen := 0
	length := len(s)
	res := ""
	for center < length {
		left, right := center, center
		for left >= 0 && s[left] == s[center] {
			left--
		}
		for right < length && s[right] == s[center] { //acafgggaggg
			right++
		}
		nextCenter := right
		for left >= 0 && right < length && s[left] == s[right] {
			left--
			right++
		}
		center = nextCenter
		midMaxLen := right - left
		if midMaxLen > maxLen {
			res = s[left+1 : right]
			maxLen = midMaxLen
		}
	}
	return res
}
