package greedy

//老师想给孩子们分发糖果，有 N个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。
//
//你需要按照以下要求，帮助老师给这些孩子分发糖果：
//
//每个孩子至少分配到 1 个糖果。
//评分更高的孩子必须比他两侧的邻位孩子获得更多的糖果。
//那么这样下来，老师至少需要准备多少颗糖果呢？
func candy(ratings []int) int {
	n := len(ratings)
	if n < 2 {
		return n
	}
	left := make([]int, n)
	for i := 1; i < n; i++ {
		if ratings[i-1] < ratings[i] {
			left[i] = left[i-1] + 1
		}
	}
	right := 0
	ans := left[n-1]
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 0
		}
		ans += max(left[i], right)
	}
	ans += n
	return ans
}
