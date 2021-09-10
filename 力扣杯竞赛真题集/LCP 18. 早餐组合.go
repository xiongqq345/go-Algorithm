package 力扣杯竞赛真题集

import "sort"

//小扣在秋日市集选择了一家早餐摊位，一维整型数组 staple 中记录了每种主食的价格，一维整型数组 drinks 中记录了每种饮料的价格。小扣的计划选择一份主食和一款饮料，且花费不超过 x 元。请返回小扣共有多少种购买方案。
//
//注意：答案需要以 1e9 + 7 (1000000007) 为底取模，如：计算初始结果为：1000000008，请返回 1
//
func breakfastNumber(staple []int, drinks []int, x int) int {
	sort.Ints(staple)
	sort.Ints(drinks)
	const mod int = 1e9 + 7
	var ans int
	for i, j := 0, len(drinks)-1; i < len(staple) && j >= 0; {
		sum := staple[i] + drinks[j]
		if sum <= x {
			i++
			ans += j + 1
			ans %= mod
		} else {
			j--
		}
	}
	return ans % mod
}
