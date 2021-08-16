package backtracking

// 整数可以被看作是其因子的乘积。
//
//例如：
//
//8 = 2 x 2 x 2;
//  = 2 x 4.
//请实现一个函数，该函数接收一个整数 n 并返回该整数所有的因子组合。
//
//注意：
//
//你可以假定 n 为永远为正数。
//因子必须大于 1 并且小于 n。
//
func getFactors(n int) [][]int {
	var ans [][]int
	var helper func([]int)
	helper = func(path []int) {
		ans = append(ans, append([]int{}, path...))
		for i := path[len(path)-2]; i*i <= path[len(path)-1]; i++ {
			if path[len(path)-1]%i == 0 {
				t := make([]int, len(path)-1)
				copy(t, path[:len(path)-1])
				t = append(t, i, path[len(path)-1]/i)
				helper(t)
			}
		}
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			helper([]int{i, n / i})
		}
	}
	return ans
}
