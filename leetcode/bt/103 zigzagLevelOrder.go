package bt

// 锯齿形（蛇形）层序遍历。
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	var ans [][]int
	for i := 0; len(q) > 0; i++ {
		ans = append(ans, []int{})
		var nq []*TreeNode
		for _, node := range q {
			ans[i] = append(ans[i], node.Val)
			if node.Left != nil {
				nq = append(nq, node.Left)
			}
			if node.Right != nil {
				nq = append(nq, node.Right)
			}
		}
		if i%2 == 1 {
			layer := ans[i]
			for k, n := 0, len(layer); k < n/2; k++ {
				layer[k], layer[n-k-1] = layer[n-k-1], layer[k]
			}
		}
		q = nq
	}
	return ans
}
