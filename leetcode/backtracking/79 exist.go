package backtracking

//给定一个m x n 二维字符网格board 和一个字符串单词word 。如果word 存在于网格中，返回 true ；否则，返回 false 。
//
//单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	used := make([][]bool, m)
	for i := range used {
		used[i] = make([]bool, n)
	}

	var helper func(int, int, int) bool
	helper = func(i, j, index int) bool {
		if index == len(word) {
			return true
		}
		if i < 0 || i >= m || j < 0 || j >= n {
			return false
		}
		if used[i][j] || board[i][j] != word[index] {
			return false
		}
		used[i][j] = true

		can := helper(i, j+1, index+1) ||
			helper(i, j-1, index+1) ||
			helper(i-1, j, index+1) ||
			helper(i+1, j, index+1)

		if can {
			return true
		}
		used[i][j] = false
		return false
	}
	for i := range board {
		for j := range board[0] {
			if board[i][j] == word[0] && helper(i, j, 0) {
				return true
			}
		}
	}
	return false
}
