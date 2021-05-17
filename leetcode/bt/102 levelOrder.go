package bt

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		res = append(res, []int{})
		var nq []*TreeNode
		for _, node := range q {
			res[i] = append(res[i], node.Val)
			if node.Left != nil {
				nq = append(nq, node.Left)
			}
			if node.Right != nil {
				nq = append(nq, node.Right)
			}
		}
		q = nq
	}
	return res
}
