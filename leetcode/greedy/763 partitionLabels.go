package greedy

// 字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。返回一个表示每个字符串片段的长度的列表。
func partitionLabels(s string) []int {
	n := len(s)
	var ans []int
	for start := 0; start < n; {
		mp := make(map[uint8]int)
		candidate := s[start]
		mp[candidate] = start
		lastPos := start
		for i := start + 1; i < n; i++ {
			if pos, ok := mp[s[i]]; ok && pos <= lastPos {
				lastPos = i
			}
			if s[i] == candidate {
				lastPos = i
			}
			mp[s[i]] = i
		}
		ans = append(ans, lastPos-start+1)
		start = lastPos + 1
	}
	return ans
}
