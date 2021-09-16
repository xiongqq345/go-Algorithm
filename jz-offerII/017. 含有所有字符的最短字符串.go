package jz_offerII

//给定两个字符串 s 和 t 。返回 s 中包含 t 的所有字符的最短子字符串。如果 s 中不存在符合条件的子字符串，则返回空字符串 "" 。
//
//如果 s 中存在多个符合条件的子字符串，返回任意一个。
//
//
//
//注意： 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/M1oyTv
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	var set [128]int
	for _, ch := range t {
		set[ch]++
	}

	cnt := len(t)
	ans := s + " "
	for i, j := 0, 0; j < len(s); j++ {
		set[s[j]]--
		if set[s[j]] >= 0 {
			cnt--
		}
		for i < j && set[s[i]] < 0 {
			set[s[i]]++
			i++
		}
		if cnt == 0 && len(ans) > len(s[i:j+1]) {
			ans = s[i : j+1]
		}
	}
	if ans == s+" " {
		return ""
	}
	return ans
}
