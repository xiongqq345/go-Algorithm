package coding

// 三步问题。有个小孩正在上楼梯，楼梯有n阶台阶，小孩一次可以上1阶、2阶或3阶。实现一种方法，计算小孩有多少种上楼梯的方式。结果可能很大，你需要对结果模1000000007。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/three-steps-problem-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func waysToStep(n int) int {
	a, b, c := 0, 1, 1
	for n > 1 {
		a, b, c = b, c, (a+b+c)%(1e9+7)
		n--
	}
	return c
}
