package _57_周赛

// 给你一个 下标从 0 开始 的整数数组 nums ，返回满足下述条件的 不同 四元组 (a, b, c, d) 的 数目 ：
//
//nums[a] + nums[b] + nums[c] == nums[d] ，且
//a < b < c < d
//
//
func countQuadruplets(a []int) int {
	n := len(a)
	var cnt int
	for target := 3; target < n; target++ {
		for i := 0; i < target-2; i++ {
			for j := i + 1; j < target-1; j++ {
				for k := j + 1; k < target; k++ {
					if a[i]+a[j]+a[k] == a[target] {
						cnt++
					}
				}
			}
		}
	}
	return cnt
}
