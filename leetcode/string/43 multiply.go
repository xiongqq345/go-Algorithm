package string

import "strconv"

// 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	ans := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		x := int(num1[i] - '0')
		for j := n - 1; j >= 0; j-- {
			y := int(num2[j] - '0')
			ans[i+j+1] += x * y
		}
	}

	var carry int
	for i := m + n - 1; i >= 0; i-- {
		sum := ans[i] + carry
		carry = sum / 10
		ans[i] = sum % 10
	}
	if ans[0] == 0 {
		ans = ans[1:]
	}
	var ansStr string
	for _, v := range ans {
		ansStr += strconv.Itoa(v)
	}
	return ansStr
}
