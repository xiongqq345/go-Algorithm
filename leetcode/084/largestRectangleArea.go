package _84

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

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
