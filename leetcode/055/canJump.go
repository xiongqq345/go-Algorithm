package _55

func canJump(nums []int) bool {
	n := len(nums)
	var most int
	for i := 0; i < n; i++ {
		if i <= most {
			if i+nums[i] > most {
				most = i + nums[i]
			}
			if most >= n-1 {
				return true
			}
		}
	}

	return false
}
