package array

func productExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	ans[0] = 1

	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	prod := 1
	for i := n - 2; i >= 0; i-- {
		prod *= nums[i+1]
		ans[i] *= prod
	}
	return ans
}
