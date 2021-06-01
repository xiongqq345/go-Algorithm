package dp

// 数组的每个下标作为一个阶梯，第 i 个阶梯对应着一个非负数的体力花费值cost[i]（下标从 0 开始）。
//
//每当你爬上一个阶梯你都要花费对应的体力值，一旦支付了相应的体力值，你就可以选择向上爬一个阶梯或者爬两个阶梯。
//
//请你找出达到楼层顶部的最低花费。在开始时，你可以选择从下标为 0 或 1 的元素作为初始阶梯。
func minCostClimbingStairs(cost []int) int {
	var pre, cur int
	for _, v := range cost {
		pre, cur = cur, min(pre+v, cur+v)
	}
	return min(pre, cur)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
