package coding

// 实现一个算法，确定一个字符串 s 的所有字符是否全都不同。
// Space:O(1)
func isUnique(str string) bool {
	var set int32
	for _, ch := range str {
		if set&(1<<(ch-'a')) != 0 {
			return false
		}
		set = set | 1<<(ch-'a')
	}
	return true
}
