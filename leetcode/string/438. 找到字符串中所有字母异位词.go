package string

//给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
//
//异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/find-all-anagrams-in-a-string
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func findAnagrams(s string, p string) []int {
	var ans []int
	var set [26]int
	for _, ch := range p {
		set[ch-97]++
	}
	var tmp [26]int
	for i := range s {
		tmp[s[i]-97]++
		if i >= len(p) {
			tmp[s[i-len(p)]-97]--
		}
		if set == tmp {
			ans = append(ans, i-len(p)+1)
		}
	}
	return ans
}
