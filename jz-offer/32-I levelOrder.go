package jz_offer

// 从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
func levelOrder(root *TreeNode) []int {
	var ans []int
	q := []*TreeNode{root}
	for i := 0; i < len(q); i++ {
		if q[i] == nil {
			continue
		}
		ans = append(ans, q[i].Val)
		q = append(q, q[i].Left, q[i].Right)
	}
	return ans
}
