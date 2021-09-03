package 力扣杯竞赛真题集

import "sort"

// 小力将 N 个零件的报价存于数组 nums。小力预算为 target，假定小力仅购买两个零件，要求购买零件的花费不超过预算，请问他有多少种采购方案。
//
//注意：答案需要以 1e9 + 7 (1000000007) 为底取模，如：计算初始结果为：1000000008，请返回 1
//
func purchasePlans(nums []int, target int) int {
	sort.Ints(nums)
	j := len(nums) - 1
	var ans int
	for i := 0; i < len(nums)-1; i++ {
		for j > i {
			if nums[i]+nums[j] <= target {
				ans += j - i
				break
			}
			j--
		}
	}
	return ans % (1e9 + 7)
}
