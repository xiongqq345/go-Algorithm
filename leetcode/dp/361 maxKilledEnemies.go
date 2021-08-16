package dp

// 想象一下炸弹人游戏，在你面前有一个二维的网格来表示地图，网格中的格子分别被以下三种符号占据：
//
//'W' 表示一堵墙
//'E' 表示一个敌人
//'0'（数字 0）表示一个空位
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/bomb-enemy
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func maxKilledEnemies(grid [][]byte) int {
	ds := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var ans int
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] != '0' {
				continue
			}
			var cnt int
			for _, d := range ds {
				for r, c := i+d[0], j+d[1]; inArea(grid, r, c); r, c = r+d[0], c+d[1] {
					if grid[r][c] == 'W' {
						break
					}
					if grid[r][c] == 'E' {
						cnt++
					}
				}
			}
			ans = max(ans, cnt)
		}
	}
	return ans
}

func inArea(grid [][]byte, i, j int) bool {
	return i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0])
}
