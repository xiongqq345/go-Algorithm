package jz_offer

// 输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
//
//序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。
func findContinuousSequence(target int) [][]int {
	var ans [][]int
	for l, r := 1, 2; l < r; {
		sum := (l + r) * (r - l + 1) / 2
		switch {
		case sum > target:
			l++
		case sum < target:
			r++
		default:
			var tmp []int
			for i := l; i <= r; i++ {
				tmp = append(tmp, i)
			}
			ans = append(ans, tmp)
			l++
		}
	}
	return ans
}
