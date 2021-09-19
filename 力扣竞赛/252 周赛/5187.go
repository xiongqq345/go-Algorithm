package _52_周赛

// 给你一个用无限二维网格表示的花园，每一个 整数坐标处都有一棵苹果树。整数坐标 (i, j) 处的苹果树有 |i| + |j| 个苹果。
//
//你将会买下正中心坐标是 (0, 0) 的一块 正方形土地 ，且每条边都与两条坐标轴之一平行。
//
//给你一个整数 neededApples ，请你返回土地的 最小周长 ，使得 至少 有 neededApples 个苹果在土地 里面或者边缘上。

func minimumPerimeter(neededApples int64) int64 {
	var num, x int64
	for x = 2; num < neededApples; x += 2 {
		num += 3 * x * x
	}
	return (x - 2) * 4
}
