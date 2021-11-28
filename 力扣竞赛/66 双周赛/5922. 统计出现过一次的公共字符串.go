package _6_双周赛

//给你两个字符串数组 words1 和 words2 ，请你返回在两个字符串数组中 都恰好出现一次 的字符串的数目。
func countWords(words1, words2 []string) (ans int) {
	cnt1 := map[string]int{}
	cnt2 := map[string]int{}
	for _, s := range words1 {
		cnt1[s]++
	}
	for _, s := range words2 {
		cnt2[s]++
	}
	for _, s := range words2 {
		if cnt1[s] == 1 && cnt2[s] == 1 {
			ans++
		}
	}
	return
}
