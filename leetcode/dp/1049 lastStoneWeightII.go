package dp

//有一堆石头，每块石头的重量都是正整数。
//
//每一回合，从中选出任意两块石头，然后将它们一起粉碎。假设石头的重量分别为x 和y，且x <= y。那么粉碎的可能结果如下：
//
//如果x == y，那么两块石头都会被完全粉碎；
//如果x != y，那么重量为x的石头将会完全粉碎，而重量为y的石头新重量为y-x。
//最后，最多只会剩下一块石头。返回此石头最小的可能重量。如果没有石头剩下，就返回 0。
func lastStoneWeightII(stones []int) int {
	var sum int
	for _, s := range stones {
		sum += s
	}
	target := sum / 2
	dp := make([]int, target+1)
	for _, v := range stones {
		for i := target; i >= v; i-- {
			dp[i] = max(dp[i], dp[i-v]+v)
		}
	}
	return sum - 2*dp[target]
}
