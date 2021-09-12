package 力扣杯竞赛

//在 n*m 大小的棋盘中，有黑白两种棋子，黑棋记作字母 "X", 白棋记作字母 "O"，空余位置记作 "."。当落下的棋子与其他相同颜色的棋子在行、列或对角线完全包围（中间不存在空白位置）另一种颜色的棋子，则可以翻转这些棋子的颜色。
//
//「力扣挑战赛」黑白翻转棋项目中，将提供给选手一个未形成可翻转棋子的棋盘残局，其状态记作 chessboard。若下一步可放置一枚黑棋，请问选手最多能翻转多少枚白棋。
//注意：
//
//若翻转白棋成黑棋后，棋盘上仍存在可以翻转的白棋，将可以 继续 翻转白棋
//输入数据保证初始棋盘状态无可以翻转的棋子且存在空余位置
func flipChess(board []string) int {
	d := []int{-1, 0, 1}
	var helper func([][]byte, int, int, int, int, int, [][2]int) int
	helper = func(tmp [][]byte, i, j, dx, dy, cnt int, warr [][2]int) int {
		if !inArea(tmp, i, j) || tmp[i][j] == '.' {
			return 0
		}
		if tmp[i][j] == 'X' {
			for _, point := range warr {
				tmp[point[0]][point[1]] = 'X'
			}

			for _, point := range warr {
				for _, dx := range d {
					for _, dy := range d {
						if dx == dy && dx == 0 {
							continue
						}
						cnt += helper(tmp, point[0]+dx, point[1]+dy, dx, dy, 0, make([][2]int, 0))
					}
				}
			}
			return cnt
		}
		if tmp[i][j] == 'O' {
			warr = append(warr, [2]int{i, j})
			cnt++
		}

		return helper(tmp, i+dx, j+dy, dx, dy, cnt, warr)
	}

	var ans int
	for i := range board {
		for j := range board[i] {
			if board[i][j] != '.' {
				continue
			}
			byteBoard := make([][]byte, len(board))
			for i := range byteBoard {
				byteBoard[i] = []byte(board[i])
			}
			byteBoard[i][j] = 'X'

			var total int
			for _, dx := range d {
				for _, dy := range d {
					if dx == dy && dx == 0 {
						continue
					}
					total += helper(byteBoard, i+dx, j+dy, dx, dy, 0, make([][2]int, 0))
				}
			}
			ans = max(ans, total)
		}
	}
	return ans
}

func inArea(board [][]byte, i, j int) bool {
	return i >= 0 && j >= 0 && i < len(board) && j < len(board[0])
}
