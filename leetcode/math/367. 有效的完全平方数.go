package math

//给定一个 正整数 num ，编写一个函数，如果 num 是一个完全平方数，则返回 true ，否则返回 false 。
//
//进阶：不要 使用任何内置的库函数，如  sqrt 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/valid-perfect-square
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func isPerfectSquare(num int) bool {
	for i := 1; i*i <= num; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}
