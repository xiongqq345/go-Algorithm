package array

// 给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。
//
//进阶：你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？
func singleNumber(nums []int) []int {
	mp := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := mp[v]; ok {
			delete(mp, v)
		} else {
			mp[v] = struct{}{}
		}
	}
	var ans []int
	for k := range mp {
		ans = append(ans, k)
	}
	return ans
}
