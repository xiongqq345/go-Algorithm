package array

// 给定两个由一些 闭区间 组成的列表，firstList 和 secondList ，其中 firstList[i] = [starti, endi] 而 secondList[j] = [startj, endj] 。每个区间列表都是成对 不相交 的，并且 已经排序 。
//
//返回这 两个区间列表的交集 。
//
//形式上，闭区间 [a, b]（其中 a <= b）表示实数 x 的集合，而 a <= x <= b 。
//
//两个闭区间的 交集 是一组实数，要么为空集，要么为闭区间。例如，[1, 3] 和 [2, 4] 的交集为 [2, 3] 。
func intervalIntersection(A [][]int, B [][]int) [][]int {
	i, j, m, n := 0, 0, len(A), len(B)
	var result [][]int
	for i < m && j < n {
		la, ra := A[i][0], A[i][1]
		lb, rb := B[j][0], B[j][1]
		// 如果有A/B有交集
		if rb >= la && ra >= lb {
			result = append(result, []int{max(la, lb), min(ra, rb)})
		}
		if rb < ra {
			j++
		} else {
			i++
		}
	}
	return result
}
