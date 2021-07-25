package dfs_bfs

// N x N 的棋盘 board 上，按从 1 到 N*N 的数字给方格编号，编号 从左下角开始，每一行交替方向。
//
//例如，一块 6 x 6 大小的棋盘，编号如下：
//
//
//r 行 c 列的棋盘，按前述方法编号，棋盘格中可能存在 “蛇” 或 “梯子”；如果 board[r][c] != -1，那个蛇或梯子的目的地将会是 board[r][c]。
//
//玩家从棋盘上的方格 1 （总是在最后一行、第一列）开始出发。
//
//每一回合，玩家需要从当前方格 x 开始出发，按下述要求前进：
//
//选定目标方格：选择从编号 x+1，x+2，x+3，x+4，x+5，或者 x+6 的方格中选出一个目标方格 s ，目标方格的编号 <= N*N。
//该选择模拟了掷骰子的情景，无论棋盘大小如何，你的目的地范围也只能处于区间 [x+1, x+6] 之间。
//传送玩家：如果目标方格 S 处存在蛇或梯子，那么玩家会传送到蛇或梯子的目的地。否则，玩家传送到目标方格 S。
//注意，玩家在每回合的前进过程中最多只能爬过蛇或梯子一次：就算目的地是另一条蛇或梯子的起点，你也不会继续移动。
//
//返回达到方格 N*N 所需的最少移动次数，如果不可能，则返回 -1。
func id2rc(id, n int) (r, c int) {
	r, c = (id-1)/n, (id-1)%n
	if r%2 == 1 {
		c = n - 1 - c
	}
	r = n - 1 - r
	return
}

func snakesAndLadders(board [][]int) int {
	n := len(board)
	vis := make([]bool, n*n+1)
	type pair struct{ id, step int }
	q := []pair{{1, 0}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for i := 1; i <= 6; i++ {
			nxt := p.id + i
			if nxt > n*n { // 超出边界
				break
			}
			r, c := id2rc(nxt, n) // 得到下一步的行列
			if board[r][c] > 0 {  // 存在蛇或梯子
				nxt = board[r][c]
			}
			if nxt == n*n { // 到达终点
				return p.step + 1
			}
			if !vis[nxt] {
				vis[nxt] = true
				q = append(q, pair{nxt, p.step + 1}) // 扩展新状态
			}
		}
	}
	return -1
}
