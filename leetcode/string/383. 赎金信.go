package string

// 为了不在赎金信中暴露字迹，从杂志上搜索各个需要的字母，组成单词来表达意思。
//
// 给你一个赎金信 (ransomnote) 字符串和一个杂志(magazine)字符串，判断 ransomnote 能不能由 magazines 里面的字符构成。
//
// 如果可以构成，返回 true ；否则返回 false 。
//
// magazine 中的每个字符只能在 ransomnote 中使用一次。
//
// 来源：力扣（leetcode）
// 链接：https://leetcode-cn.com/problems/ransom-note
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func canconstruct(ransomnote string, magazine string) bool {
	if len(ransomnote) > len(magazine) {
		return false
	}
	var set [26]int
	for _, v := range magazine {
		set[v-'a']++
	}
	for _, v := range ransomnote {
		if set[v-'a'] <= 0 {
			return false
		}
		set[v-'a']--
	}
	return true
}
