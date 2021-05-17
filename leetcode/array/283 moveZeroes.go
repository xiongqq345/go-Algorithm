package array

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
func moveZeroes(nums []int) {
	i := 0
	for _, v := range nums {
		if v != 0 {
			nums[i] = v
			i++
		}
	}
	for i < len(nums) {
		nums[i] = 0
		i++
	}
}
