package _48

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	for _, v := range nums {
		v = (v - 1) % n
		nums[v] += n
	}
	var res []int
	for i, v := range nums {
		if v <= n {
			res = append(res, i+1)
		}
	}
	return res
}
