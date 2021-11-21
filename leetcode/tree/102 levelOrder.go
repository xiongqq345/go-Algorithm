package tree

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ans [][]int
	q := []*TreeNode{root}
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
		q = nq
	}
	return ans
}
