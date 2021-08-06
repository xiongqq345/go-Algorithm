package dp

// 给定一个未排序的整数数组，找到最长递增子序列的个数。
func findNumberOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//两个状态变量
	//dp为到达第i个位置时的最长子序列长度
	dp := make([]int, len(nums))
	//count表示到达第I个位置时的最长子序列长度有几种情况
	count := make([]int, len(nums))
	for i := range nums {
		dp[i], count[i] = 1, 1
	}
	maxx := 1 //定义初始化最大值
	res := 0  //定义最长递增子序列的个数
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				//如果dp[i] < dp[j]+1 表明在0-i-1中，出现了
				//新的最大值，则将最大值更新
				if dp[i] < dp[j]+1 {
					//更新最大值
					dp[i] = dp[j] + 1
					//此时count[i]=count[j]；以nums[i]结尾
					//的最长递增子序列的组合方式就等于nums[i]目前的组合方式
					count[i] = count[j]
				} else if dp[i] == dp[j]+1 { //当相等时，又发现一种情况，
					//将现在的情况与j时的组合方式相加
					count[i] += count[j]
				}
			}
			maxx = max(maxx, dp[i])
		}
	}
	//  遍历遍历dp[i]:找到最大长度，
	// 然后将结果加上相应的组合长度
	for i := 0; i < len(count); i++ {
		if dp[i] == maxx {
			res += count[i]
		}
	}
	return res
}
