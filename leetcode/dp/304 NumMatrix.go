package dp

// 给定一个二维矩阵 matrix，以下类型的多个请求：
//
//计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1, col1) ，右下角为 (row2, col2) 。
//实现 NumMatrix 类：
//
//NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
//int sumRegion(int row1, int col1, int row2, int col2) 返回左上角 (row1, col1) 、右下角 (row2, col2) 的子矩阵的元素总和。

type NumMatrix struct {
	mat, pre [][]int
	m, n     int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	P := make([][]int, m+1)
	for i := range P {
		P[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			P[i][j] = P[i-1][j] + P[i][j-1] - P[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return NumMatrix{
		mat: matrix,
		pre: P,
		m:   m,
		n:   n,
	}
}

func (nm *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return nm.get(row2+1, col2+1) + nm.get(row1, col1) - nm.get(row2+1, col1) - nm.get(row1, col2+1)
}

func (nm *NumMatrix) get(x, y int) int {
	x = max(min(x, nm.m), 0)
	y = max(min(y, nm.n), 0)
	return nm.pre[x][y]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
