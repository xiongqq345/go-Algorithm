package _5_双周赛

//给你一个在 XY 平面上的 width x height 的网格图，左下角 的格子为 (0, 0) ，右上角 的格子为 (width - 1, height - 1) 。网格图中相邻格子为四个基本方向之一（"North"，"East"，"South" 和 "West"）。一个机器人 初始 在格子 (0, 0) ，方向为 "East" 。
//
//机器人可以根据指令移动指定的 步数 。每一步，它可以执行以下操作。
//
//沿着当前方向尝试 往前一步 。
//如果机器人下一步将到达的格子 超出了边界 ，机器人会 逆时针 转 90 度，然后再尝试往前一步。
//如果机器人完成了指令要求的移动步数，它将停止移动并等待下一个指令。
//
//请你实现 Robot 类：
//
//Robot(int width, int height) 初始化一个 width x height 的网格图，机器人初始在 (0, 0) ，方向朝 "East" 。
//void move(int num) 给机器人下达前进 num 步的指令。
//int[] getPos() 返回机器人当前所处的格子位置，用一个长度为 2 的数组 [x, y] 表示。
//String getDir() 返回当前机器人的朝向，为 "North" ，"East" ，"South" 或者 "West" 。

type Robot struct {
	w, h   int
	dx, dy int
	offset int
}

func Constructor(width int, height int) Robot {
	return Robot{
		w:  width - 1,
		h:  height - 1,
		dx: 1,
	}
}

func (r *Robot) Move(num int) {
	c := r.w*2 + r.h*2
	rem := num % c
	r.offset += rem
	r.offset %= c
	r.dx, r.dy = 0, 0
	switch {
	case r.offset == 0:
		r.dy = -1
	case r.offset <= r.w:
		r.dx = 1
	case r.offset <= c/2:
		r.dy = 1
	case r.offset <= r.w+r.h+r.w:
		r.dx = -1
	case r.offset <= c:
		r.dy = -1
	}
}

func (r *Robot) GetPos() []int {
	switch {
	case r.dx == 1:
		return []int{r.offset, 0}
	case r.dx == -1:
		return []int{r.w - (r.offset - r.w - r.h), r.h}
	case r.dy == 1:
		return []int{r.w, r.offset - r.w}
	case r.dy == -1:
		if r.offset == 0 {
			return []int{0, 0}
		}
		return []int{0, 2*(r.w+r.h) - r.offset}
	}
	return nil
}

func (r *Robot) GetDir() string {
	switch {
	case r.dy == -1:
		return "South"
	case r.dy == 1:
		return "North"
	case r.dx == 1:
		return "East"
	case r.dx == -1:
		return "West"
	}
	return ""
}

/**
 * Your Robot object will be instantiated and called as such:
 * obj := Constructor(width, height);
 * obj.Move(num);
 * param_2 := obj.GetPos();
 * param_3 := obj.GetDir();
 */
