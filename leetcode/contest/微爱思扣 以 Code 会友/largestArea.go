package 微爱思扣_以_Code_会友

// 「以扣会友」线下活动所在场地由若干主题空间与走廊组成，场地的地图记作由一维字符串型数组 grid，字符串中仅包含 "0"～"5" 这 6 个字符。地图上每一个字符代表面积为 1 的区域，其中 "0" 表示走廊，其他字符表示主题空间。相同且连续（连续指上、下、左、右四个方向连接）的字符组成同一个主题空间。
//
//假如整个 grid 区域的外侧均为走廊。请问，不与走廊直接相邻的主题空间的最大面积是多少？如果不存在这样的空间请返回 0。

func largestArea(grid []string) int {
	m, n := len(grid), len(grid[0])
	arr := make([][]int, m)
	for i := range arr {
		arr[i] = make([]int, n)
	}

	for i := range arr {
		for j := range arr[0] {
			arr[i][j] = int(grid[i][j])
		}
	}

}
