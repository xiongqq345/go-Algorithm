package jz_offer

// 输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。如果有多对数字的和等于s，则输出任意一对即可。
func twoSum(nums []int, target int) []int {
	p1, p2 := 0, len(nums)-1
	for p1 < p2 {
		sum := nums[p1] + nums[p2]
		switch {
		case sum > target:
			p2--
		case sum < target:
			p1++
		default:
			return []int{nums[p1], nums[p2]}
		}
	}
	return nil
}
