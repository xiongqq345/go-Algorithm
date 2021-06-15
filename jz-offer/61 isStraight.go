package jz_offer

func isStraight(nums []int) bool {
	mp := make(map[int]int)

	min, max := 14, 0
	for _, num := range nums {
		if num != 0 {
			if num > max {
				max = num
			}
			if num < min {
				min = num
			}
			mp[num]++
		}
	}
	for _, v := range mp {
		if v > 1 {
			return false
		}
	}
	return max-min < 5
}
