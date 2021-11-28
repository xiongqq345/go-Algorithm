package string

//给你两个字符串 s 和 goal ，只要我们可以通过交换 s 中的两个字母得到与 goal 相等的结果，就返回 true ；否则返回 false 。
//
//交换字母的定义是：取两个下标 i 和 j （下标从 0 开始）且满足 i != j ，接着交换 s[i] 和 s[j] 处的字符。
//
//例如，在 "abcd" 中交换下标 0 和下标 2 的元素可以生成 "cbad" 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/buddy-strings
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func buddyStrings(s, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	var set1, set2 [26]int
	var cnt int
	var hasSame bool
	for i := range s {
		set1[s[i]-'a']++
		set2[goal[i]-'a']++
		if set1[s[i]-'a'] == 2 {
			hasSame = true
		}
		if s[i] != goal[i] {
			cnt++
		}
	}
	if set1 != set2 {
		return false
	}

	return cnt == 2 || cnt == 0 && hasSame
}
