package _2_双周赛

//给你一个 数字 字符串数组 nums 和一个 数字 字符串 target ，请你返回 nums[i] + nums[j] （两个字符串连接）结果等于 target 的下标 (i, j) （需满足 i != j）的数目。
//
//
func numOfPairs(nums []string, target string) int {
	var ans int
	for i := range nums {
		for j := range nums {
			if i == j {
				continue
			}
			if nums[i]+nums[j] == target {
				ans++
			}
		}
	}
	return ans
}
