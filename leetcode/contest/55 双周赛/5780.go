package _5_双周赛

func canBeIncreasing(nums []int) bool {
	for i := range nums {
		tmp := append(nums[:i], nums[i+1:]...)
		var flag bool
		for j := 1; j < len(tmp); j++ {
			if tmp[j] <= tmp[j-1] {
				flag = true
				break
			}
		}
		if !flag {
			return true
		}
	}
	return false
}
