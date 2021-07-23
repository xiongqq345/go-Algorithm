package string

// 给定 S 和 T 两个字符串，当它们分别被输入到空白的文本编辑器后，判断二者是否相等，并返回结果。 # 代表退格字符。
//
//注意：如果对空文本输入退格字符，文本继续为空。

func backspaceCompare(S string, T string) bool {
	return Com(S) == Com(T)
}

func Com(s string) string {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] != '#' {
			stack = append(stack, s[i])
		} else if len(stack) != 0 {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}
