package jz_offerII

// 给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
func singleNumber(nums []int) int {
	var a, b int
	for _, num := range nums {
		b = (b ^ num) &^ a
		a = (a ^ num) &^ b
	}
	return b
}
