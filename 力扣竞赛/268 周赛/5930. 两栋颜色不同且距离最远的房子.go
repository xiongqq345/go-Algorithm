package _68_周赛

//街上有 n 栋房子整齐地排成一列，每栋房子都粉刷上了漂亮的颜色。给你一个下标从 0 开始且长度为 n 的整数数组 colors ，其中 colors[i] 表示第  i 栋房子的颜色。
//
//返回 两栋 颜色 不同 房子之间的 最大 距离。
//
//第 i 栋房子和第 j 栋房子之间的距离是 abs(i - j) ，其中 abs(x) 是 x 的绝对值。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/two-furthest-houses-with-different-colors
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func maxDistance(colors []int) int {
	var ans int
	for i := range colors {
		for j := len(colors) - 1; j > i; j-- {
			if colors[j] != colors[i] {
				ans = max(ans, j-i)
				break
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
