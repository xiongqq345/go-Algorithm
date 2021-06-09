package jz_offer

// 一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。
//
//答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
func numWays(n int) int {
	pre, cur := 1, 1
	for ; n > 0; n-- {
		pre, cur = cur, (pre+cur)%(1e9+7)
	}
	return pre
}
