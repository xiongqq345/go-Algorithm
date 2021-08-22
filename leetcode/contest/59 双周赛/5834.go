package _9_双周赛

// 有一个特殊打字机，它由一个 圆盘 和一个 指针 组成， 圆盘上标有小写英文字母 'a' 到 'z'。只有 当指针指向某个字母时，它才能被键入。指针 初始时 指向字符 'a' 。
//
//
//每一秒钟，你可以执行以下操作之一：
//
//将指针 顺时针 或者 逆时针 移动一个字符。
//键入指针 当前 指向的字符。
//给你一个字符串 word ，请你返回键入 word 所表示单词的 最少 秒数 。
func minTimeToType(word string) int {
	ans := len(word)
	cur := 'a'
	for _, v := range word {
		t := abs(int(v - cur))
		if t >= 0 && t < 13 {
			ans += t
		} else {
			ans += 26 - t
		}
		cur = v
	}
	return ans
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
