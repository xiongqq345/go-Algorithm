package jz_offerII

//给定一个二维矩阵 matrix，以下类型的多个请求：
//
//计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1, col1) ，右下角为 (row2, col2) 。
//实现 NumMatrix 类：
//
//NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
//int sumRegion(int row1, int col1, int row2, int col2) 返回左上角 (row1, col1) 、右下角 (row2, col2) 的子矩阵的元素总和。
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/O4NDxx
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
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
