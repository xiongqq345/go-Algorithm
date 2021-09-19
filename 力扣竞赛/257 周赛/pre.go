package _57_周赛

const mod int = 1e9 + 7

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

type heapInt []int

//Less  小于就是小跟堆，大于号就是大根堆
func (h *heapInt) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *heapInt) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *heapInt) Len() int           { return len(*h) }
func (h *heapInt) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *heapInt) Pop() interface{} {
	t := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return t
}
func (h *heapInt) Peek() int {
	return (*h)[0]
}

func a() {
	//初始化map
	//mp := make(map[int]int)
	//for _, t := range tasks {
	//	mp[t]++
	//}

	//背包问题
	//for _, coin := range coins {
	//	for i := coin; i <= amount; i++ {
	//		dp[i] = min(dp[i], dp[i-coin]+1)
	//	}
	//}
}
