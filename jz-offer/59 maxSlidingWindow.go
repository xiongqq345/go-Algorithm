package jz_offer

// 给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。
//
//示例:
//
//输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
//输出: [3,3,5,5,6,7]
//解释:
//
//  滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//

func maxSlidingWindow(nums []int, k int) []int {
	var ans, deque []int
	for i := 0; i < len(nums); i++ {
		for len(deque) > 0 && deque[len(deque)-1] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, nums[i])
		if i >= k && deque[0] == nums[i-k] {
			deque = deque[1:]
		}
		if i >= k-1 {
			ans = append(ans, deque[0])
		}
	}
	return ans
}
