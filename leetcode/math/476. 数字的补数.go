package math

//对整数的二进制表示取反（0 变 1 ，1 变 0）后，再转换为十进制表示，可以得到这个整数的补数。
//
//例如，整数 5 的二进制表示是 "101" ，取反后得到 "010" ，再转回十进制表示得到补数 2 。
//给你一个整数 num ，输出它的补数。
//
func findComplement(num int) int {
	highBit := 0
	for i := 1; i <= 30; i++ {
		if num < 1<<i {
			break
		}
		highBit = i
	}
	mask := 1<<(highBit+1) - 1
	return num ^ mask
}
