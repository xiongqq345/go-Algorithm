package leetcode

func removeElement(nums []int, val int) int {
	lens := len(nums)
	count := 0
	for i := 0; i < lens; i++ {
		if nums[i-count] == val {
			nums = append(nums[:i-count], nums[i+1-count:lens-count]...)
			nums = append(nums, val)
			count++
		}
	}
	return lens - count
}

func removeElement2(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
	return j
}
