package jz_offerII

//一个专业的小偷，计划偷窃一个环形街道上沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
//
//给定一个代表每个房屋存放金额的非负整数数组 nums ，请计算 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。
//
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(_rob(nums[1:]), _rob(nums[:len(nums)-1]))
}

func _rob(nums []int) int {
	var pre, cur int
	for _, v := range nums {
		pre, cur = cur, max(pre+v, cur)
	}
	return cur
}
