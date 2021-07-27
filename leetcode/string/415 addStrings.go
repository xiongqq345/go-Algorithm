package string

// 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
//
//
//
//提示：
//
//num1 和num2 的长度都小于 5100
//num1 和num2 都只包含数字 0-9
//num1 和num2 都不包含任何前导零
//你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式
//
func addStrings(num1 string, num2 string) string {
	var ans []byte
	var carry, sum uint8
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		var v1, v2 uint8
		if i >= 0 {
			v1 = num1[i] - '0'
		}
		if j >= 0 {
			v2 = num2[j] - '0'
		}
		sum = v1 + v2 + carry
		carry, sum = (sum)/10, (sum)%10
		ans = append(ans, sum+'0')

	}
	if carry > 0 {
		ans = append(ans, carry+'0')
	}
	for i := range ans[:len(ans)/2] {
		ans[i], ans[len(ans)-i-1] = ans[len(ans)-i-1], ans[i]
	}
	return string(ans)
}
