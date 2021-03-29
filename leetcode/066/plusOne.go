package _66

func plusOne(digits []int) []int {
	var addOne bool
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] == 10 {
			digits[i] = 0
			addOne = true
		} else {
			addOne = false
			break
		}
	}

	if addOne {
		digits = append([]int{1}, digits...)
	}
	return digits
}
