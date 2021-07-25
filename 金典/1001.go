package coding

// 给定两个排序后的数组 A 和 B，其中 A 的末端有足够的缓冲空间容纳 B。 编写一个方法，将 B 合并入 A 并排序。
//
//初始化 A 和 B 的元素数量分别为 m 和 n。
//
func merge(A []int, m int, B []int, n int) {
	p := m + n - 1
	i, j := m-1, n-1
	for i >= 0 && j >= 0 {
		if A[i] < B[j] {
			A[p] = B[j]
			j--
		} else {
			A[p] = A[i]
			i--
		}
		p--
	}
	if i <= 0 {
		copy(A, B[:j+1])
	}
}
