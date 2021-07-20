package coding

//二进制数转字符串。给定一个介于0和1之间的实数（如0.72），类型为double，打印它的二进制表达式。如果该数字无法精确地用32位以内的二进制表示，则打印“ERROR”。
func printBin(num float64) string {
	res := "0."
	for num != float64(0) {
		num *= 2
		if num >= 1 {
			res += "1"
			num -= 1.0
		} else {
			res += "0"
		}
		if len(res) > 32 {
			return "ERROR"
		}
	}
	return res
}
