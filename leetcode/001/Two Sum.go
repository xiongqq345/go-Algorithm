package leetcode

func twoSum(nums []int, target int) []int {
	maps := make(map[int]int)
	nes := make([]int, 0)

	for key, val := range nums {
		ant := target - val
		_, st := maps[ant]
		if st {
			nes = append(nes, maps[ant], key)
			return nes
		}
		maps[val] = key
	}
	return nes
}
