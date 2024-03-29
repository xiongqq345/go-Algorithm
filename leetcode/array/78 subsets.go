package array

//给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
//
//解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
func subsets(nums []int) [][]int {
	var ans [][]int
	for mask := 0; mask < 1<<len(nums); mask++ {
		var set []int
		for i, v := range nums {
			if mask>>i&1 == 1 {
				set = append(set, v)
			}
		}
		ans = append(ans, set)
	}
	return ans
}
