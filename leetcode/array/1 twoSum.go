package array

func twoSum(nums []int, target int) []int {
	ht := make(map[int]int)
	for i, v := range nums {
		if val, ok := ht[v]; ok {
			return []int{val, i}
		}
		ht[target-v] = i
	}
	return nil
}
