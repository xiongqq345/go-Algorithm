package list_stack

// 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
//求在该柱状图中，能够勾勒出来的矩形的最大面积。
func largestRectangleArea(heights []int) int {
	var stack []int
	var ans int
	heights = append(append([]int{0}, heights...), 0)
	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left, right := stack[len(stack)-1], i
			ans = max(ans, (right-left-1)*heights[cur])
		}
		stack = append(stack, i)
	}
	return ans
}
