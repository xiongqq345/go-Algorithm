package _66_周赛

func countVowels(word string) int64 {
	var ans int64
	for i := range word {
		for j := i; j <= len(word); j++ {
			ans += count(word[i:j])
		}
	}
	return ans
}

func count(s string) (cnt int64) {
	for _, v := range s {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			cnt++
		}
	}
	return cnt
}
