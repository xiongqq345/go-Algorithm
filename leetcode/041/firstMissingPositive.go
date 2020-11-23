package leetcode

func firstMissingPositive(nums []int) int {
	m := make(map[int]struct{})
	for _, v := range nums {
		m[v] = struct{}{}
	}

	for i := 0; i < len(nums)+1; i++ {
		if _, ok := m[i+1]; !ok {
			return i + 1
		}
	}
	return len(nums) + 1
}

func firstMissingPositive2(nums []int) int {
	n := len(nums)
	for i, v := range nums {
		if v <= 0 {
			nums[i] = n + 1
		}
	}
	for _, v := range nums {
		if abs(v) <= n {
			nums[abs(v)-1] = abs(nums[abs(v)-1]) * -1
		}
	}
	for i, v := range nums {
		if v > 0 {
			return i + 1
		}
	}
	return n + 1
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
