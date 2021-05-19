package string

// 实现一个算法，确定一个字符串 s 的所有字符是否全都不同。
func isUnique(astr string) bool {
	m := make(map[int32]struct{})
	for _, c := range astr {
		if _, ok := m[c]; ok {
			return false
		}
		m[c] = struct{}{}
	}
	return true
}
