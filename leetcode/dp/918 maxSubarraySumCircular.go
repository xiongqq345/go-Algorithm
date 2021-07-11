package dp

import "math"

// 给定一个由整数数组 A 表示的环形数组 C，求 C 的非空子数组的最大可能和。
//
//在此处，环形数组意味着数组的末端将会与开头相连呈环状。（形式上，当0 <= i < A.length 时 C[i] = A[i]，且当 i >= 0 时 C[i+A.length] = C[i]）
//
//此外，子数组最多只能包含固定缓冲区 A 中的每个元素一次。（形式上，对于子数组 C[i], C[i+1], ..., C[j]，不存在 i <= k1, k2 <= j 其中 k1 % A.length = k2 % A.length）
//
func maxSubarraySumCircular(nums []int) int {
	var total int
	dpMax, maxSubArr := 0, nums[0]
	dpMin, minSubArr := 0, nums[0]
	maxV := math.MinInt32
	for _, v := range nums {
		total += v
		maxV = max(maxV, v)
		dpMax = max(dpMax+v, v)
		maxSubArr = max(maxSubArr, dpMax)

		dpMin = min(dpMin+v, v)
		minSubArr = min(minSubArr, dpMin)
	}

	if maxV < 0 {
		return maxV
	}
	return max(maxSubArr, total-minSubArr)
}
