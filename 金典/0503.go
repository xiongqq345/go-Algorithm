package coding

// 给定一个32位整数 num，你可以将一个数位从0变为1。请编写一个程序，找出你能够获得的最长的一串1的长度。
func reverseBits(num int) int {
	if num < 0 {
		num += 1 << 32
	}
	ans, n1, n2 := 0, 0, 0
	for {
		if num&1 == 1 {
			n1++
			if n1 >= 32 {
				break
			}
		} else {
			ans = max(ans, n1+n2+1)
			if num == 0 {
				break
			}
			n1, n2 = 0, n1
		}
		num >>= 1
	}
	ans = max(ans, n1+n2)
	return ans
}
