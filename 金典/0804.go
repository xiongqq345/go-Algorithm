package coding

// 幂集。编写一种方法，返回某集合的所有子集。集合中不包含重复的元素。
//
//说明：解集不能包含重复的子集。
func subsets(nums []int) [][]int {
	n := len(nums)
	var ans [][]int
	for mask := 0; mask < 1<<n; mask++ {
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
