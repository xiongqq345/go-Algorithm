package math

//和谐数组是指一个数组里元素的最大值和最小值之间的差别 正好是 1 。
//
//现在，给你一个整数数组 nums ，请你在所有可能的子序列中找到最长的和谐子序列的长度。
//
//数组的子序列是一个由数组派生出来的序列，它可以通过删除一些元素或不删除元素、且不改变其余元素的顺序而得到。
//
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-harmonious-subsequence
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func findHS(nums []int) int {
	mp := make(map[int]int)
	for _, v := range nums {
		mp[v]++
	}
	var ans int
	for k,v := range mp {
		if mp[k+1] > 0 {
			ans = max(ans, v+mp[k+1])
		}
	}
	return ans
}
