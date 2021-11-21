package tree

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	ans := 0
	for len(q) > 0 {
		var nq []*TreeNode
		for i := 0; i < len(q); i++ {
			if q[i].Left != nil {
				nq = append(nq, q[i].Left)
			}
			if q[i].Right != nil {
				nq = append(nq, q[i].Right)
			}
		}
		q = nq
		ans++
	}
	return ans
}
