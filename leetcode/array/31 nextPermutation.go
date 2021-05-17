package array

// 下一个排列
func nextPermutation(nums []int) {
	n := len(nums)
	l, r := n-2, n-1
	for l >= 0 && nums[l] >= nums[l+1] {
		l--
	}
	if l >= 0 { // 不是最后一个排列
		for r >= 0 && nums[r] <= nums[l] {
			r--
		}
		nums[l], nums[r] = nums[r], nums[l]
	}

	reverse(nums[l+1:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-i-1] = a[n-i-1], a[i]
	}
}
