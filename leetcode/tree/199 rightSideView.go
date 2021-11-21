package tree

// 右侧视图
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	q := []*TreeNode{root}
	for len(q) > 0 {
		n := len(q)
		res = append(res, q[n-1].Val)
		for i := 0; i < n; i++ {
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		q = q[n:]
	}
	return res
}
