package _7_双周赛

//给你一个炸弹列表。一个炸弹的 爆炸范围 定义为以炸弹为圆心的一个圆。
//
//炸弹用一个下标从 0 开始的二维整数数组 bombs 表示，其中 bombs[i] = [xi, yi, ri] 。xi 和 yi 表示第 i 个炸弹的 X 和 Y 坐标，ri 表示爆炸范围的 半径 。
//你需要选择引爆 一个 炸弹。当这个炸弹被引爆时，所有 在它爆炸范围内的炸弹都会被引爆，这些炸弹会进一步将它们爆炸范围内的其他炸弹引爆。
//
//给你数组 bombs ，请你返回在引爆 一个 炸弹的前提下，最多 能引爆的炸弹数目。
func maximumDetonation(bombs [][]int) int {
	n := len(bombs)
	var ans int
	var dfs func(int, *int)
	vis := make([]bool, n)
	dfs = func(idx int, num *int) {
		vis[idx] = true
		*num++
		ans = max(ans, *num)
		if *num >= n {
			return
		}
		for i := 0; i < n; i++ {
			if vis[i] {
				continue
			}
			if in(bombs[idx], bombs[i]) {
				dfs(i, num)
			}
		}
	}

	for i := range bombs {
		if ans == n {
			return n
		}
		num := new(int)
		dfs(i, num)
		for i := range vis {
			vis[i] = false
		}
	}
	return ans
}

func in(a, b []int) bool {
	return (a[0]-b[0])*(a[0]-b[0])+(a[1]-b[1])*(a[1]-b[1]) <= a[2]*a[2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
