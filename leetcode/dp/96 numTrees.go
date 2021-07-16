package dp

// 给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？

//G(n)=f(1)+f(2)+f(3)+f(4)+...+f(n)
//
//当 i 为根节点时，其左子树节点个数为 i-1 个，右子树节点为 n-i，则
//f(i) = G(i-1)*G(n-i)
//
//综合两个公式可以得到 卡特兰数 公式
//G(n) = G(0)*G(n-1)+G(1)*(n-2)+...+G(n-1)*G(0)

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i < n+1; i++ {
		for j := 1; j < i+1; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

func numTrees2(n int) int {
	c := 1
	for i := 0; i < n; i++ {
		//c *= (4*i + 2) / (i + 2)
		c = c * (4*i + 2) / (i + 2)
	}
	return c
}
