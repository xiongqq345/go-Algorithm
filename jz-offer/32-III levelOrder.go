package jz_offer

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
			for k, n := 0, len(layer)-1; k < len(layer)/2; k++ {
				layer[k], layer[n-k] = layer[n-k], layer[k]
			}
		}
		q = nq
	}
	return res
}
