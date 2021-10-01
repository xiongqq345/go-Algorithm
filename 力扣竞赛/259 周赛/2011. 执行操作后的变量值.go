package _59_周赛

//存在一种仅支持 4 种操作和 1 个变量 X 的编程语言：
//
//++X 和 X++ 使变量 X 的值 加 1
//--X 和 X-- 使变量 X 的值 减 1
//最初，X 的值是 0
//
//给你一个字符串数组 operations ，这是由操作组成的一个列表，返回执行所有操作后， X 的 最终值 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/final-value-of-variable-after-performing-operations
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func finalValueAfterOperations(operations []string) int {
	var x int
	for _, v := range operations {
		if v[1] == '+' {
			x++
		} else {
			x--
		}
	}
	return x
}
