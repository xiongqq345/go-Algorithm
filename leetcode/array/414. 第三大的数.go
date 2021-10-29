package array

//给你一个非空数组，返回此数组中 第三大的数 。如果不存在，则返回数组中最大的数。
func thirdMax(nums []int) int {
	var a, b, c *int
	for _, num := range nums {
		num := num
		if a == nil || num > *a {
			a, b, c = &num, a, b
		} else if *a > num && (b == nil || num > *b) {
			b, c = &num, b
		} else if b != nil && *b > num && (c == nil || num > *c) {
			c = &num
		}
	}
	if c == nil {
		return *a
	}
	return *c
}
