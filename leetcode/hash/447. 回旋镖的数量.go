package hash

// 给定平面上 n 对 互不相同 的点 points ，其中 points[i] = [xi, yi] 。回旋镖 是由点 (i, j, k) 表示的元组 ，其中 i 和 j 之间的距离和 i 和 k 之间的距离相等（需要考虑元组的顺序）。
//
//返回平面上所有回旋镖的数量。
//
func numberOfBoomerangs(points [][]int) int {
	var ans int
	for _, p := range points {
		mp := make(map[int]int)
		for _, q := range points {
			dis := (p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1])
			mp[dis]++
		}
		for _, cnt := range mp {
			ans += cnt * (cnt - 1)
		}
	}
	return ans
}
