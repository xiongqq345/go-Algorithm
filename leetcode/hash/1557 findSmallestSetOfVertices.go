package hash

// 给你一个 有向无环图 ， n 个节点编号为 0 到 n-1 ，以及一个边数组 edges ，其中 edges[i] = [fromi, toi] 表示一条从点  fromi 到点 toi 的有向边。
//
//找到最小的点集使得从这些点出发能到达图中所有点。题目保证解存在且唯一。
//
//你可以以任意顺序返回这些节点编号。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/minimum-number-of-vertices-to-reach-all-nodes
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func findSmallestSetOfVertices(n int, edges [][]int) []int {
	set := make(map[int]bool)
	for _, v := range edges {
		set[v[1]] = true
	}
	var ans []int
	for i := 0; i < n; i++ {
		if !set[i] {
			ans = append(ans, i)
		}
	}
	return ans
}
