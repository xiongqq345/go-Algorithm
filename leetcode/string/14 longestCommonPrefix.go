package string

//编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 ""。
func longestCommonPrefix(strs []string) string {
	for index := 0; index < len(strs[0]); index++ {
		c := strs[0][index]
		for i := 1; i < len(strs); i++ {
			if index == len(strs[i]) || strs[i][index] != c {
				return strs[0][:index]
			}
		}
	}
	return strs[0]
}
