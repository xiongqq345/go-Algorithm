package _53_周赛

// 给你一个字符串 s ，下标从 0 开始 ，且长度为偶数 n 。字符串 恰好 由 n / 2 个开括号 '[' 和 n / 2 个闭括号 ']' 组成。
//
//只有能满足下述所有条件的字符串才能称为 平衡字符串 ：
//
//字符串是一个空字符串，或者
//字符串可以记作 AB ，其中 A 和 B 都是 平衡字符串 ，或者
//字符串可以写成 [C] ，其中 C 是一个 平衡字符串 。
//你可以交换 任意 两个下标所对应的括号 任意 次数。
//
//返回使 s 变成 平衡字符串 所需要的 最小 交换次数。
//
func minSwaps(s string) int {
	var cnt, mincnt int
	for _, ch := range s {
		if ch == '[' {
			cnt++
		} else {
			cnt--
			mincnt = min(mincnt, cnt)
		}
	}
	if mincnt == 0 {
		return 0
	}
	return (-mincnt + 1) / 2
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
