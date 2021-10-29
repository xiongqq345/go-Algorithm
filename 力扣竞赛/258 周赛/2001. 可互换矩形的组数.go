package _58_周赛

//用一个下标从 0 开始的二维整数数组 rectangles 来表示 n 个矩形，其中 rectangles[i] = [widthi, heighti] 表示第 i 个矩形的宽度和高度。
//
//如果两个矩形 i 和 j（i < j）的宽高比相同，则认为这两个矩形 可互换 。更规范的说法是，两个矩形满足 widthi/heighti == widthj/heightj（使用实数除法而非整数除法），则认为这两个矩形 可互换 。
//
//计算并返回 rectangles 中有多少对 可互换 矩形。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/number-of-pairs-of-interchangeable-rectangles
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func interchangeableRectangles(rectangles [][]int) int64 {
	mp := make(map[[2]int]int64)
	for _, v := range rectangles {
		w, h := v[0], v[1]
		g := gcd(w, h)
		mp[[2]int{w / g, h / g}]++
	}
	var ans int64
	for _, v := range mp {
		ans += v * (v - 1) / 2
	}
	return ans
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
