package design

type NumMatrix struct {
	prefix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	pre := make([][]int, len(matrix)+1)
	pre[0] = make([]int, len(matrix[0])+1)
	for i := range matrix {
		var sum int
		pre[i+1] = make([]int, len(matrix[0])+1)
		for j := range matrix[i] {
			sum += matrix[i][j]
			pre[i+1][j+1] = sum + pre[i][j+1]
		}
	}

	return NumMatrix{
		prefix: pre,
	}
}

func (nm *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return nm.prefix[row2+1][col2+1] - nm.prefix[row2+1][col1] - nm.prefix[row1][col2+1] + nm.prefix[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
