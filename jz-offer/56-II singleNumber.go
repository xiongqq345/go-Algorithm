package jz_offer

// 在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。
func singleNumber(nums []int) int {
	var ones, twos int
	for _, num := range nums {
		ones = ones ^ num & ^twos
		twos = twos ^ num & ^ones
	}
	return ones
}
