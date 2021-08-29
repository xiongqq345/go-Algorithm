package array

// 给你一个正整数数组 arr ，请你计算所有可能的奇数长度子数组的和。
//
//子数组 定义为原数组中的一个连续子序列。
//
//请你返回 arr 中 所有奇数长度子数组的和 。
func sumOddLengthSubarrays(arr []int) int {
	sum, n := 0, len(arr)
	for i := range arr {
		leftCount, rightCount := i, n-i-1
		leftOdd := (leftCount + 1) / 2
		rightOdd := (rightCount + 1) / 2
		leftEven := leftCount/2 + 1
		rightEven := rightCount/2 + 1
		sum += arr[i] * (leftOdd*rightOdd + leftEven*rightEven)
	}
	return sum
}
