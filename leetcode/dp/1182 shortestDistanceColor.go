package dp

// 给你一个数组 colors，里面有  1、2、 3 三种颜色。
//
//我们需要在 colors 上进行一些查询操作 queries，其中每个待查项都由两个整数 i 和 c 组成。
//
//现在请你帮忙设计一个算法，查找从索引 i 到具有目标颜色 c 的元素之间的最短距离。
//
//如果不存在解决方案，请返回 -1。
//
func shortestDistanceColor(colors []int, queries [][]int) []int {
	n := len(colors)
	l, r := make([][3]int, n), make([][3]int, n)
	var ans []int
	for i := 0; i < n; i++ {
		l[i] = [3]int{-1, -1, -1}
		r[i] = [3]int{-1, -1, -1}
	}

	// get left
	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			if i-1 >= 0 && l[i-1][j] != -1 {
				l[i][j] = l[i-1][j] + 1
			}
		}
		l[i][colors[i]-1] = 0
	}

	// get right
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < 3; j++ {
			if i+1 < n && r[i+1][j] != -1 {
				r[i][j] = r[i+1][j] + 1
			}
		}
		r[i][colors[i]-1] = 0
	}

	for _, q := range queries {
		idx, c := q[0], q[1]-1
		d := 1 << 31
		if l[idx][c] != -1 {
			d = min(d, l[idx][c])
		}
		if r[idx][c] != -1 {
			d = min(d, r[idx][c])
		}
		if d == 1<<31 {
			ans = append(ans, -1)
		} else {
			ans = append(ans, d)
		}
	}
	return ans
}
