package jz_offer

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		var nq []*TreeNode
		res = append(res, []int{})
		for j := 0; j < len(q); j++ {
			res[i] = append(res[i], q[j].Val)
			if q[j].Left != nil {
				nq = append(nq, q[j].Left)
			}
			if q[j].Right != nil {
				nq = append(nq, q[j].Right)
			}
		}
		q = nq
	}
	return res
}
