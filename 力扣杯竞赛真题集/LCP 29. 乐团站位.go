package 力扣杯竞赛真题集

// 某乐团的演出场地可视作 num * num 的二维矩阵 grid（左上角坐标为 [0,0])，每个位置站有一位成员。乐团共有 9 种乐器，乐器编号为 1~9，每位成员持有 1 个乐器。
//
//为保证声乐混合效果，成员站位规则为：自 grid 左上角开始顺时针螺旋形向内循环以 1，2，...，9 循环重复排列。

func orchestraLayout(num int, xPos int, yPos int) int {
	grid := make([][]int, num)
	for i := range grid {
		grid[i] = make([]int, num)
	}
	l, t, r, b := 0, 0, num-1, num-1
	for n, c := 1, 1; n <= num*num; {
		for i := l; i <= r; i++ {
			if t == xPos && i == yPos {
				return c
			}
			grid[t][i] = c
			c++
			n++
			if c == 10 {
				c = 1
			}
		}
		t++
		for i := t; i <= b; i++ {
			if i == xPos && r == yPos {
				return c
			}
			grid[i][r] = c
			c++
			n++
			if c == 10 {
				c = 1
			}
		}
		r--
		for i := r; i >= l; i-- {
			if b == xPos && i == yPos {
				return c
			}
			grid[b][i] = c
			c++
			n++
			if c == 10 {
				c = 1
			}
		}
		b--
		for i := b; i >= t; i-- {
			if i == xPos && l == yPos {
				return c
			}
			grid[i][l] = c
			c++
			n++
			if c == 10 {
				c = 1
			}
		}
		l++
	}
	return 0
}
