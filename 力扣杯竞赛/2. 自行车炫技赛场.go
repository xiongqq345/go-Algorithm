package 力扣杯竞赛

import (
	"sort"
)

func bicycleYard(position []int, terrain [][]int, obstacle [][]int) [][]int {
	ds := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	//m, n := len(terrain), len(terrain[0])
	var ans [][]int
	set := make(map[[2]int]bool)
	set[[2]int{position[0], position[1]}] = true

	mp := make(map[[4]int]int)
	var helper func(int, int, int)
	helper = func(speed, i, j int) {
		if speed <= 0 {
			return
		}
		if speed == 1 && !set[[2]int{i, j}] {
			ans = append(ans, []int{i, j})
			set[[2]int{i, j}] = true
		}

		for _, d := range ds {
			x, y := i+d[0], j+d[1]
			if inArea(terrain, x, y) {
				helper(speed+terrain[i][j]-terrain[x][y]-obstacle[x][y], x, y)
			}
		}
	}
	helper(1, position[0], position[1])
	sort.Slice(ans, func(i, j int) bool {
		if ans[i][0] != ans[j][0] {
			return ans[i][0] < ans[j][0]
		}
		return ans[i][1] < ans[j][1]
	})
	return ans
}

func inArea(grid [][]int, i, j int) bool {
	return i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0])
}
