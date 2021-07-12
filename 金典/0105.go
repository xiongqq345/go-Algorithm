package coding

//字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。
func oneEditAway(first string, second string) bool {
	if len(first) > len(second) {
		first, second = second, first
	}
	if len(second)-len(first) > 1 {
		return false
	}

	for i := range first {
		if first[i] != second[i] {
			if first[i:] == second[i+1:] || first[i+1:] == second[i+1:] {
				return true
			}
			return false
		}
	}
	return true
}
