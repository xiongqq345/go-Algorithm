package leetcode

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length < 2 {
		return length
	}
	element := nums[0]
	p1, p2 := 0, 1
	for p2 < length {
		fastV := nums[p2]
		if fastV == element {
			p2++
			continue
		} else {
			element = fastV
			nums[p1+1] = element
			p1++
		}
	}
	return p1 + 1
}

func removeDuplicates2(nums []int) int {
	cursor := 0
	for i := 1; i < len(nums); i++ {
		if nums[cursor] != nums[i] {
			cursor++
			nums[cursor] = nums[i]
		}
	}
	return cursor + 1
}
