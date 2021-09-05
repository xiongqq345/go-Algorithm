package 力扣杯竞赛真题集

// 某乐团的演出场地可视作 num * num 的二维矩阵 grid（左上角坐标为 [0,0])，每个位置站有一位成员。乐团共有 9 种乐器，乐器编号为 1~9，每位成员持有 1 个乐器。
//
//为保证声乐混合效果，成员站位规则为：自 grid 左上角开始顺时针螺旋形向内循环以 1，2，...，9 循环重复排列。

func orchestraLayout(num int, xPos int, yPos int) int {
	x, y, n := xPos, yPos, num
	if x <= y {
		k := min(x, n-1-y)
		return (4*k*(n-k)+1+(x+y-k*2)-1)%9 + 1
	}
	kp := min(y, n-1-x) + 1
	return (4*kp*(n-kp)+1-(x+y-(kp-1)*2)-1)%9 + 1
}
