package dp

// 给你一个整数数组 nums ，你可以对它进行一些操作。
//
//每次操作中，选择任意一个 nums[i] ，删除它并获得 nums[i] 的点数。之后，你必须删除 所有 等于 nums[i] - 1 和 nums[i] + 1 的元素。
//
//开始你拥有 0 个点数。返回你能通过这些操作获得的最大点数。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/delete-and-earn
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func deleteAndEarn(nums []int) int {
	maxVal := 0
	for _, v := range nums {
		maxVal = max(maxVal, v)
	}

	bucket := make([]int, maxVal)
	for _, v := range nums {
		bucket[v-1]++
	}
	var pre, cur int
	for i, v := range bucket {
		pre, cur = cur, max(cur, pre+(i+1)*v)
	}
	return cur
}
