package array

// 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点(i,ai) 。在坐标内画 n 条垂直线，垂直线 i的两个端点分别为(i,ai) 和 (i, 0) 。找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	var maxArea int
	for i < j {
		area := (j - i) * min(height[i], height[j])
		maxArea = max(maxArea, area)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxArea
}
