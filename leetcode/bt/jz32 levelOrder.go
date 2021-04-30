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
