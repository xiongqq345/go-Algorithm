package 力扣杯竞赛

// 小扣注意到秋日市集上有一个创作黑白方格画的摊位。摊主给每个顾客提供一个固定在墙上的白色画板，画板不能转动。画板上有 n * n 的网格。绘画规则为，小扣可以选择任意多行以及任意多列的格子涂成黑色（选择的整行、整列均需涂成黑色），所选行数、列数均可为 0。
//
//小扣希望最终的成品上需要有 k 个黑色格子，请返回小扣共有多少种涂色方案。
//
//注意：两个方案中任意一个相同位置的格子颜色不同，就视为不同的方案。
func C(base, up int) int {
	a := 1
	for i := 0; i < up; i++ {
		a *= base - i
	}

	b := 1

	for i := 1; i <= up; i++ {
		b *= i
	}
	return a / b
}

func paintingPlan(n int, k int) int {
	if k == n*n {
		return 1
	}

	ret := 0
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			if i*n+j*n-i*j == k {
				ret += C(n, i) * C(n, j)
			}
		}
	}
	return ret
}
