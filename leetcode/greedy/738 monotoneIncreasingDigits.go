package greedy

import "strconv"

//给定一个非负整数 N，找出小于或等于 N 的最大的整数，同时这个整数需要满足其各个位数上的数字是单调递增。
//
//（当且仅当每个相邻位数上的数字 x 和 y 满足 x <= y 时，我们称这个整数是单调递增的。）
//n   = 1234321
//res = 1233999
func monotoneIncreasingDigits(n int) int {
	str := strconv.Itoa(n)
	idx := -1
	arr := make([]byte, len(str))
	for i := 0; i < len(str)-1; i++ {
		arr[i] = str[i]
		if str[i] > str[i+1] {
			arr[i] = str[i] - 1
			idx = i
			break
		}
	}
	if idx != -1 {
		for i := idx; i < len(str)-1; i++ {
			arr[i+1] = '9'
		}
	}
	res, _ := strconv.Atoi(string(arr))
	return res
}
