package _2_双周赛

//给你一个下标从 0 开始的一维整数数组 original 和两个整数 m 和  n 。你需要使用 original 中 所有 元素创建一个 m 行 n 列的二维数组。
//
//original 中下标从 0 到 n - 1 （都 包含 ）的元素构成二维数组的第一行，下标从 n 到 2 * n - 1 （都 包含 ）的元素构成二维数组的第二行，依此类推。
//
//请你根据上述过程返回一个 m x n 的二维数组。如果无法构成这样的二维数组，请你返回一个空的二维数组。

func construct2DArray(original []int, m int, n int) [][]int {
	k := len(original)
	if k != m*n {
		return nil
	}
	var ans [][]int
	for i := 0; i < k; i += n {
		ans = append(ans, original[i:i+n])
	}
	return ans
}
