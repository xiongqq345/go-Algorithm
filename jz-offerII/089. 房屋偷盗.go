package jz_offerII

//一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响小偷偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//
//给定一个代表每个房屋存放金额的非负整数数组 nums ，请计算 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
//
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/Gu0c2T
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func rob(nums []int) int {
	var pre, cur int
	for _, v := range nums {
		pre, cur = cur, max(pre+v, cur)
	}
	return cur
}
