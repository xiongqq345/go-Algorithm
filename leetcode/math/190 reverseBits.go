package math

func reverseBits(num uint32) uint32 {
	var ans uint32
	for i := 0; i < 32 && num > 0; i++ {
		ans |= num & 1 << (31 - i)
		num >>= 1
	}
	return ans
}
