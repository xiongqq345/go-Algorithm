package jz_offerII

//给定非负整数数组 heights ，数组中的数字用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
//求在该柱状图中，能够勾勒出来的矩形的最大面积。
//
//
func largestRectangleArea(heights []int) int {
	var ans int
	var stack []int
	heights = append(append([]int{0}, heights...), 0)
	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			cur := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			l := stack[len(stack)-1]
			ans = max(ans, (i-l-1)*cur)
		}
		stack = append(stack, i)
	}
	return ans
}
