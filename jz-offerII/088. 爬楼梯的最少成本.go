package jz_offerII

//数组的每个下标作为一个阶梯，第 i 个阶梯对应着一个非负数的体力花费值 cost[i]（下标从 0 开始）。
//
//每当爬上一个阶梯都要花费对应的体力值，一旦支付了相应的体力值，就可以选择向上爬一个阶梯或者爬两个阶梯。
//
//请找出达到楼层顶部的最低花费。在开始时，你可以选择从下标为 0 或 1 的元素作为初始阶梯。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/GzCJIP
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func minCostClimbingStairs(cost []int) int {
	var pre, cur int
	for _, v := range cost {
		pre, cur = cur, min(pre, cur)+v
	}
	return min(pre, cur)
}
