package leetcode

func spiralOrder(matrix [][]int) []int {
	n := len(matrix)
	m := len(matrix[0])
	var ring int
	if m < n {
		ring = (m + 1) / 2
	} else {
		ring = (n + 1) / 2
	}
	for r := 0; r < ring; r++ {
		for n := 0; n < (m-r*2+n-r*2)*2; n++ {

		}
	}
}
