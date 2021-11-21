package tree

//给定一个二叉树，返回其节点值自底向上的层序遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		res = append(res, []int{})
		var nq []*TreeNode
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

	for i, n := 0, len(res); i < n/2; i++ {
		res[i], res[n-i-1] = res[n-i-1], res[i]
	}
	return res
}
