package bt

func levelOrder(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for i := 0; i < len(q); i++ {
		if q[i] == nil {
			continue
		}
		res = append(res, q[i].Val)
		q = append(q, q[i].Left, q[i].Right)
	}
	return res
}

func levelOrder(root *TreeNode) [][]int {
	q := []*TreeNode{root}
	res := make([][]int, 0)
	if root == nil {
		return res
	}
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
		if i%2 == 1 {
			layer := res[i]
			for k, n := 0, len(layer); k < n/2; k++ {
				layer[k], layer[n-k-1] = layer[n-k-1], layer[k]
			}
		}
		q = nq
	}
	return res
}
