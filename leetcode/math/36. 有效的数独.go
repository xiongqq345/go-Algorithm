package math

import "strconv"

//请你判断一个 9x9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。
//
//数字 1-9 在每一行只能出现一次。
//数字 1-9 在每一列只能出现一次。
//数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
//数独部分空格内已填入了数字，空白格用 '.' 表示。
//
//注意：
//
//一个有效的数独（部分已被填充）不一定是可解的。
//只需要根据以上规则，验证已经填入的数字是否有效即可。
//
func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if !isValid(board, 0, i, 8, i) {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if !isValid(board, i, 0, i, 8) {
			return false
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !isValid(board, 3*i, 3*j, 3*i+2, 3*j+2) {
				return false
			}
		}
	}
	return true
}

func isValid(board [][]byte, x1, y1, x2, y2 int) bool {
	presence := make([]int, 9)
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			if board[i][j] == '.' {
				continue
			}
			c, _ := strconv.Atoi(string(board[i][j]))
			if presence[c-1] == 1 {
				return false
			} else {
				presence[c-1] = 1
			}
		}
	}
	return true
}
