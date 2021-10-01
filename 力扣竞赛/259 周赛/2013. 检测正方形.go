package _59_周赛

//给你一个在 X-Y 平面上的点构成的数据流。设计一个满足下述要求的算法：
//
//添加 一个在数据流中的新点到某个数据结构中。可以添加 重复 的点，并会视作不同的点进行处理。
//给你一个查询点，请你从数据结构中选出三个点，使这三个点和查询点一同构成一个 面积为正 的 轴对齐正方形 ，统计 满足该要求的方案数目。
//轴对齐正方形 是一个正方形，除四条边长度相同外，还满足每条边都与 x-轴 或 y-轴 平行或垂直。
//
//实现 DetectSquares 类：
//
//DetectSquares() 使用空数据结构初始化对象
//void add(int[] point) 向数据结构添加一个新的点 point = [x, y]
//int count(int[] point) 统计按上述方式与点 point = [x, y] 共同构造 轴对齐正方形 的方案数。

type DetectSquares struct {
	rows [1001][]int
	cols [1001]map[int]int
}

func Constructor() DetectSquares {
	var cols [1001]map[int]int
	for i := range cols {
		cols[i] = make(map[int]int)
	}
	return DetectSquares{
		rows: [1001][]int{},
		cols: cols,
	}
}

func (ds *DetectSquares) Add(point []int) {
	x, y := point[0], point[1]
	ds.rows[y] = append(ds.rows[y], x)
	ds.cols[x][y]++
}

func (ds *DetectSquares) Count(point []int) int {
	var cnt int
	x, y := point[0], point[1]
	for _, x2 := range ds.rows[y] {
		if x2 != x {
			d := abs(x2 - x)
			cnt += ds.cols[x2][y-d] * ds.cols[x][y-d]
			cnt += ds.cols[x2][y+d] * ds.cols[x][y+d]
		}
	}
	return cnt
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */
