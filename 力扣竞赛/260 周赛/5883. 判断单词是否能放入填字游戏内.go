package _60_周赛

//给你一个 m x n 的矩阵 board ，它代表一个填字游戏 当前 的状态。填字游戏格子中包含小写英文字母（已填入的单词），表示 空 格的 ' ' 和表示 障碍 格子的 '#' 。
//
//如果满足以下条件，那么我们可以 水平 （从左到右 或者 从右到左）或 竖直 （从上到下 或者 从下到上）填入一个单词：
//
//该单词不占据任何 '#' 对应的格子。
//每个字母对应的格子要么是 ' ' （空格）要么与 board 中已有字母 匹配 。
//如果单词是 水平 放置的，那么该单词左边和右边 相邻 格子不能为 ' ' 或小写英文字母。
//如果单词是 竖直 放置的，那么该单词上边和下边 相邻 格子不能为 ' ' 或小写英文字母。
//给你一个字符串 word ，如果 word 可以被放入 board 中，请你返回 true ，否则请返回 false 。
//

func placeWordInCrossword(board [][]byte, word string) bool {
	ds := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var helper func(int, int, int, int, int) bool
	helper = func(i, j, di, dj, pos int) bool {
		if pos == len(word) {
			return !inArea(board, i, j) || board[i][j] == '#'
		}
		if !inArea(board, i, j) || board[i][j] == '#' || (board[i][j] != ' ' && board[i][j] != word[pos]) {
			return false
		}
		if pos == 0 {
			if di != 0 && inArea(board, i-di, j) && board[i-di][j] != '#' && board[i-di][j] != word[pos] {
				return false
			}
			if dj != 0 && inArea(board, i, j-dj) && board[i][j-dj] != '#' && board[i][j-dj] != word[pos] {
				return false
			}
		}

		return helper(i+di, j+dj, di, dj, pos+1)
	}

	for i := range board {
		for j := range board[i] {
			for _, d := range ds {
				if helper(i, j, d[0], d[1], 0) {
					return true
				}
			}
		}
	}
	return false
}

func inArea(grid [][]byte, i, j int) bool {
	return i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0])
}
