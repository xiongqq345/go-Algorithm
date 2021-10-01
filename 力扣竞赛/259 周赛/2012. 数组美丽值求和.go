package _59_周赛

//给你一个下标从 0 开始的整数数组 nums 。对于每个下标 i（1 <= i <= nums.length - 2），nums[i] 的 美丽值 等于：
//
//2，对于所有 0 <= j < i 且 i < k <= nums.length - 1 ，满足 nums[j] < nums[i] < nums[k]
//1，如果满足 nums[i - 1] < nums[i] < nums[i + 1] ，且不满足前面的条件
//0，如果上述条件全部不满足
//返回符合 1 <= i <= nums.length - 2 的所有 nums[i] 的 美丽值的总和 。
//
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sum-of-beauty-in-the-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func sumOfBeauties(a []int) int {
	n := len(a)
	sufMin := make([]int, n)
	sufMin[n-1] = a[n-1]
	for i := n - 2; i >= 0; i-- {
		sufMin[i] = min(sufMin[i+1], a[i])
	}

	var ans int
	maximum := a[0]
	for i := 1; i < n-1; i++ {
		if a[i] > maximum && a[i] < sufMin[i+1] {
			ans++
		}
		maximum = max(maximum, a[i])
		if a[i-1] < a[i] && a[i] < a[i+1] {
			ans++
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
