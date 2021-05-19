package array

// 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
func rotate2(nums []int, k int) {
	n := len(nums)
	k %= n
	temp := make([]int, k)
	copy(temp, nums[n-k:])
	copy(nums[k:], nums[:n-k])
	copy(nums[:k], temp)
}

//func rotate3(nums []int, k int) {
//	k %= len(nums)
//	copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
//}
//
//func rotate3(nums []int, k int) {
//	k %= len(nums)
//	reverse(nums)
//	reverse(nums[:k])
//	reverse(nums[k:])
//}
//
//func reverse(a []int) {
//	for i, n := 0, len(a); i < n/2; i++ {
//		a[i], a[n-i-1] = a[n-i-1], a[i]
//	}
//}
