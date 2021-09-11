package 力扣杯竞赛

// 秋日市集上，魔术师邀请小扣与他互动。魔术师的道具为分别写有数字 1~N 的 N 张卡牌，然后请小扣思考一个 N 张卡牌的排列 target。
//
//魔术师的目标是找到一个数字 k（k >= 1），使得初始排列顺序为 1~N 的卡牌经过特殊的洗牌方式最终变成小扣所想的排列 target，特殊的洗牌方式为：
//
//第一步，魔术师将当前位于 偶数位置 的卡牌（下标自 1 开始），保持 当前排列顺序 放在位于 奇数位置 的卡牌之前。例如：将当前排列 [1,2,3,4,5] 位于偶数位置的 [2,4] 置于奇数位置的 [1,3,5] 前，排列变为 [2,4,1,3,5]；
//第二步，若当前卡牌数量小于等于 k，则魔术师按排列顺序取走全部卡牌；若当前卡牌数量大于 k，则取走前 k 张卡牌，剩余卡牌继续重复这两个步骤，直至所有卡牌全部被取走；
//卡牌按照魔术师取走顺序构成的新排列为「魔术取数排列」，请返回是否存在这个数字 k 使得「魔术取数排列」恰好就是 target，从而让小扣感到大吃一惊。

func isMagic(target []int) bool {
	nums := make([]int, len(target))
	for i := range nums {
		nums[i] = i + 1
	}
	nums = od(nums)
	var k int
	for k < len(nums) && nums[k] == target[k] {
		k++
	}
	if k == 0 {
		return false
	}
	for {
		if k >= len(nums) {
			return isSame(nums, target, len(nums))
		}
		if !isSame(nums, target, k) {
			return false
		}
		nums = nums[k:]
		target = target[k:]
		nums = od(nums)
	}
}

func isSame(nums, target []int, end int) bool {
	for i := 0; i < end; i++ {
		if nums[i] != target[i] {
			return false
		}
	}
	return true
}

func od(nums []int) []int {
	var idx int
	res := make([]int, len(nums))
	for i := 1; i < len(nums); i += 2 {
		res[idx] = nums[i]
		idx++
	}
	for i := 0; i < len(nums); i += 2 {
		res[idx] = nums[i]
		idx++
	}
	return res
}
