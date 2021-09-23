package jz_offerII

//请根据每日 气温 列表 temperatures ，重新生成一个列表，要求其对应位置的输出为：要想观测到更高的气温，至少需要等待的天数。如果气温在这之后都不会升高，请在该位置用 0 来代替。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/iIQa4I
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func dailyTemperatures(temperatures []int) []int {
	var st []int
	ans := make([]int, len(temperatures))
	for i, t := range temperatures {
		for len(st) > 0 && temperatures[st[len(st)-1]] < t {
			ans[st[len(st)-1]] = i - st[len(st)-1]
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	return ans
}
