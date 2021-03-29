package _80

func removeDuplicates(nums []int) int {
	i, j := 1, 2
	for j < len(nums) {
		if nums[j] != nums[i-1] {
			i++
			nums[i] = nums[j]
		}
		j++
	}
	return i + 1
}
