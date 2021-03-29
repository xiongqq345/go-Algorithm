package _78

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	for mask := 0; mask < 1<<len(nums); mask++ {
		set := make([]int, 0)
		for i, v := range nums {
			if mask>>i&1 > 0 {
				set = append(set, v)
			}
		}
		res = append(res, set)
	}
	return res
}
