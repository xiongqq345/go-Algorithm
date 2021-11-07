package _66_周赛

func countVowelSubstrings(word string) int {
	var ans int
	for i := range word {
		for j := i; j <= len(word); j++ {
			if is(word[i:j]) {
				ans++
			}
		}
	}
	return ans
}

func is(s string) bool {
	var a, e, i, o, u bool
	for _, v := range s {
		if v != 'a' && v != 'e' && v != 'i' && v != 'o' && v != 'u' {
			return false
		}
		switch v {
		case 'a':
			a = true
		case 'e':
			e = true
		case 'i':
			i = true
		case 'o':
			o = true
		case 'u':
			u = true
		}
	}
	return a && e && i && o && u
}
