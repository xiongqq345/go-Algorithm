package array

//给定一个二进制数组， 计算其中最大连续 1 的个数。
//
func findMaxConsecutiveOnes(nums []int) int {
	var cnt, ans int
	for _, num := range nums {
		if num == 1 {
			cnt++
		} else {
			ans = max(ans, cnt)
			cnt = 0
		}
	}
	ans = max(ans, cnt)
	return ans
}
