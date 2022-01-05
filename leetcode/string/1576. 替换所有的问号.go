package string

//给你一个仅包含小写英文字母和 '?' 字符的字符串 s，请你将所有的 '?' 转换为若干小写字母，使最终的字符串不包含任何 连续重复 的字符。
//
//注意：你 不能 修改非 '?' 字符。
//
//题目测试用例保证 除 '?' 字符 之外，不存在连续重复的字符。
//
//在完成所有转换（可能无需转换）后返回最终的字符串。如果有多个解决方案，请返回其中任何一个。可以证明，在给定的约束条件下，答案总是存在的。
func modifyString(s string) string {
	res := []byte(s)
	n := len(res)
	for i, ch := range res {
		if ch != '?' {
			continue
		}
		for b := byte('a'); b <= 'c'; b++ {
			if !(i > 0 && res[i-1] == b || i < n-1 && res[i+1] == b) {
				res[i] = b
				break
			}
		}
	}
	return string(res)
}
