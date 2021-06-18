package math

//编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为汉明重量）。
func hammingWeight(num uint32) int {
	var ans int
	for num > 0 {
		num &= num - 1
		ans++
	}
	return ans
}
