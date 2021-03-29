package _283

func moveZeroes(nums []int) {
	var index int
	for _, v := range nums {
		if v != 0 {
			nums[index] = v
			index++
		}
	}
	for index < len(nums) {
		nums[index] = 0
		index++
	}
}
