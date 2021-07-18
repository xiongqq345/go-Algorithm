package dp

// 给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			row[j] += row[j-1]
		}
	}
	return row
}
