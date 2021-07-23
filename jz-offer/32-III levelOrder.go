package jz_offer

func levelOrder(root *TreeNode) [][]int {
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
			for k, n := 0, len(layer)-1; k < len(layer)/2; k++ {
				layer[k], layer[n-k] = layer[n-k], layer[k]
			}
		}
		q = nq
	}
	return ans
}
