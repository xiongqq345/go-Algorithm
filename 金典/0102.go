package coding

// 给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。
func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	var sum1, sum2 int32
	for i := range s1 {
		sum1 += 1 << (s1[i] - 'a')
		sum2 += 1 << (s2[i] - 'a')
	}
	return sum1 == sum2
}
