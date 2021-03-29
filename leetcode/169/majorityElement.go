package _169

func majorityElement(nums []int) int {
	var candidate, count int
	for _, v := range nums {
		if count == 0 {
			candidate, count = v, 1
			continue
		}
		if v == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
